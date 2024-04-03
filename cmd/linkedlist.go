// Package cmd /*
package cmd

import (
	linkedlist "github.com/psoladoye/datastructures/adts/linkedlist"
	"log"

	"github.com/spf13/cobra"
)

// linkedlistCmd represents the linkedlist command
var linkedlistCmd = &cobra.Command{
	Use:   "linkedlist",
	Short: "A linked list is a data structure of linearly linked elements",
	Long:  `This linked list contains pointers for head and tail`,
	Run: func(cmd *cobra.Command, args []string) {
		list := linkedlist.LinkedList(&linkedlist.SinglyLinkedList{})

		list.Append(0)
		list.Prepend(20)
		list.Prepend(1)
		list.Prepend(11)
		list.Prepend(9)
		list.Prepend(34)
		list.Print()

		list.Append(200)
		list.Print()

		log.Println(list.Head().GetData())
		log.Println(list.Tail().GetData())
	},
}

func init() {
	rootCmd.AddCommand(linkedlistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkedlistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkedlistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
