package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/mishamyrt/mysh/v1/pkg/hosts"
	"github.com/mishamyrt/mysh/v1/pkg/paths"
	"github.com/mishamyrt/mysh/v1/pkg/remotes"
	"github.com/mishamyrt/mysh/v1/pkg/ssh"
)

// GitCommit refers to commit hash at the moment of build
var GitCommit string

// Version of Mysh
var Version string

// BuildTime of Mysh
var BuildTime string

func getRemote(args []string) {
	if len(args) < 3 {
		fmt.Println("Can't load repository. URL not provided. Usage:")
		fmt.Println("\tmysh get <url>")
		return
	}
	namespaceName, err := remotes.GetConfig(args[2])
	if err != nil {
		return
	}
	fmt.Printf("'%s' repository successfully added\n", namespaceName)
	err = hosts.BuildCompletionList()
	if err != nil {
		fmt.Println("Could not update shell completion")
	}
}

func printCommands(commands [][]string, writer *tabwriter.Writer) {
	for _, command := range commands {
		fmt.Fprintf(writer, "\t%s\t%s\n", command[0], command[1])
	}
}

func help() {
	fmt.Printf("Mysh is a tool for improving SSH user experience\n\n")
	fmt.Println("Usage:")
	fmt.Printf("\t mysh <command or host> [arguments]\n\n")
	fmt.Println("The commands are:")
	w := tabwriter.NewWriter(os.Stdout, 9, 1, 1, ' ', 0)
	printCommands([][]string{
		{"copy", "copy remote file"},
		{"get", "add repository and download hosts from it"},
		{"help", "print this message and exit"},
		{"hosts", "display all hosts"},
		{"namespaces", "display all namespaces"},
		{"remotes", "display all added remote repositories"},
		{"show", "display host information"},
		{"update", "refresh hosts from added remote repositories"},
		{"version", "print Mysh version"},
	}, w)
	w.Flush()
}

func updateRemotes() {
	remotes.UpdateRemotes()
	err := hosts.BuildCompletionList()
	if err != nil {
		fmt.Println("Could not update shell completion")
	}
}

func printRemotes() {
	remotes := remotes.GetRemotes()
	if len(remotes.Remotes) == 0 {
		fmt.Println("No remote namespaces were found")
		return
	}
	fmt.Println("Remote namespaces:")
	for namespace := range remotes.Remotes {
		fmt.Printf("- %s\n", namespace)
	}
}

func printNamespaces() {
	namespaces := hosts.GetNamespaces()
	if len(namespaces) == 0 {
		fmt.Println("No namespaces were found")
		return
	}
	fmt.Println("Namespaces:")
	for _, namespace := range namespaces {
		fmt.Printf("- %s\n", namespace)
	}
}

func printHosts() {
	hosts, _ := hosts.GetHosts(true)
	if len(hosts) == 0 {
		fmt.Println("No hosts were found")
		return
	}
	fmt.Println("Hosts:")
	for host := range hosts {
		fmt.Printf("- %s\n", host)
	}
}

func version() {
	w := tabwriter.NewWriter(os.Stdout, 14, 1, 1, ' ', 0)
	fmt.Fprintf(w, "Version:\t%s\n", Version)
	fmt.Fprintf(w, "GitCommit:\t%s\n", GitCommit)
	fmt.Fprintf(w, "Built:\t%s\n", BuildTime)
	w.Flush()
}

func connect(args []string) {
	var privateKey string
	var hostString string
	for i := 0; i < len(args); i++ {
		if args[i] == "-i" {
			privateKey = args[i+1]
			i++
		} else {
			hostString = args[i]
		}
	}
	host, _ := hosts.MatchHost(hostString, false)
	if len(privateKey) > 0 {
		host.Key = privateKey
	}
	command, err := ssh.BuildSSHCommand(host)
	if err != nil {
		panic(err)
	}
	fmt.Println(command)
}

func showHost(args []string) {
	if len(args) < 3 {
		fmt.Println("Host not provided. Usage:")
		fmt.Println("\tmysh show <host>")
		return
	}
	host, err := hosts.MatchHost(os.Args[2], true)
	if err != nil {
		fmt.Println("Host not found")
		return
	}
	fmt.Println("Host:", host.Host)
	if len(host.User) > 0 {
		fmt.Println("User:", host.User)
	}
	if len(host.Port) > 0 {
		fmt.Println("Port:", host.Port)
	}
}

func getRSyncFile(filePath string) (string, error) {
	if strings.Contains(filePath, ":") {
		return ssh.BuildRSyncPath(hosts.MatchRemoteFile(filePath))
	}
	return filePath, nil
}

func copyFile(args []string) {
	usage := "\tmysh copy [args] <source> <destination>"
	pathOffset := 2
	rsyncArgs := "av8h"
	if args[2][0:1] == "-" {
		pathOffset++
		rsyncArgs += args[2][1:len(args[2])]
	}
	var source string
	var target string
	if len(args) < 3 {
		fmt.Println("Source not provided. Usage:")
		fmt.Println(usage)
		return
	} else if len(args) < 4 {
		fmt.Println("Target not provided. Usage:")
		fmt.Println(usage)
		return
	}
	source, err := getRSyncFile(args[pathOffset])
	if err != nil {
		fmt.Println("Source not provided. Usage:")
		fmt.Println(usage)
	}
	target, err = getRSyncFile(args[pathOffset+1])
	if err != nil {
		fmt.Println("Target not provided. Usage:")
		fmt.Println(usage)
	}
	fmt.Printf("rsync -%s %s %s\n", rsyncArgs, source, target)
}

func main() {
	err := paths.PreapreEnvironment()
	if err != nil {
		fmt.Println("Сan't initialize configuration files")
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		help()
		return
	}
	switch os.Args[1] {
	case "get":
		getRemote(os.Args)
	case "update":
		updateRemotes()
	case "help":
		help()
	case "remotes":
		printRemotes()
	case "namespaces":
		printNamespaces()
	case "hosts":
		printHosts()
	case "show":
		showHost(os.Args)
	case "copy":
		copyFile(os.Args)
	case "version":
		version()
	default:
		connect(os.Args)
	}
}
