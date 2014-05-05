package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
)

type Command struct {
    Name        string
    Value       string
    SubCommands []Command
    Before      func()
}

var commands = []Command{
    Command{
        Name:  "dotfiles",
        Value: "git clone git://github.com/j-carvalho/dotfiles ~/.dotfiles",
        Before: func() {
            cmd := exec.Command("rm", "-rf", "~/.dotfiles")

            err := cmd.Run()

            if err != nil {
                log.Fatal(err)
            }
        },
        SubCommands: []Command{
            Command{
                Name:  "vimrc",
                Value: "ln -s .dotfiles/vimrc ~/.vimrc",
                Before: func() {
                    cmd := exec.Command("rm", "~/.vimrc")

                    err := cmd.Run()

                    if err != nil {
                        log.Fatal(err)
                    }
                },
            },
            Command{
                Name:  "vim",
                Value: "ln -s .dotfiles/vim ~/.vim",
                Before: func() {
                    cmd := exec.Command("rm", "~/.vimrc")

                    err := cmd.Run()

                    if err != nil {
                        log.Fatal(err)
                    }
                },
            },
            Command{
                Name:  "gitconfig",
                Value: "ln -s .dotfiles/gitconfig ~/.gitconfig",
                Before: func() {
                    cmd := exec.Command("rm", "~/.gitconfig")

                    err := cmd.Run()

                    if err != nil {
                        log.Fatal(err)
                    }
                },
            },
            Command{
                Name:  "gitconfig",
                Value: "ln -s .dotfiles/zsh ~/.zsh",
                Before: func() {
                    cmd := exec.Command("rm", "-rf", "~/.zsh")

                    err := cmd.Run()

                    if err != nil {
                        log.Fatal(err)
                    }
                },
            },
            Command{
                Name:  "gitconfig",
                Value: "ln -s .dotfiles/zshrc ~/.zshrc",
                Before: func() {
                    cmd := exec.Command("rm", "~/.zshrc")

                    err := cmd.Run()

                    if err != nil {
                        log.Fatal(err)
                    }
                },
            },
        },
    },
    Command{
        Name:  "Vundle",
        Value: "git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim",
        Before: func() {
            cmd := exec.Command("rm", "-rf", "~/.vim/bundle/Vundle.vim")

            err := cmd.Run()

            if err != nil {
                log.Fatal(err)
            }
        },
        SubCommands: []Command{
            Command{
                Name:  "Plugin Install",
                Value: "vim +PluginInstall +qall",
            },
        },
    },
    Command{
        Name:  "oh-my-zsh",
        Value: "git clone git://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh",
        Before: func() {
            cmd := exec.Command("rm", "-rf", "~/.oh-my-zsh")

            err := cmd.Run()

            if err != nil {
                log.Fatal(err)
            }
        },
    },
}

func (command Command) Run() {
    fmt.Printf("Initializing %s\n", command.Name)

    if command.Before != nil {
        command.Before()
    }

    args := strings.Split(command.Value, " ")

    var cmd *exec.Cmd

    if len(args) > 1 {
        cmd = exec.Command(args[0], args[1:len(args)]...)
    } else {
        cmd = exec.Command(args[0])
    }

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }

    for _, subcommand := range command.SubCommands {
        subcommand.Run()
    }

    fmt.Printf("Finished %s\n", command.Name)
}

func main() {
    for _, command := range commands {
        command.Run()
    }
}
