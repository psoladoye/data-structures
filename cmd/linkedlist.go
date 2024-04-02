/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	linkedlist "github.com/psoladoye/datastructures/adts"
	"github.com/spf13/cobra"
)

// linkedlistCmd represents the linkedlist command
var linkedlistCmd = &cobra.Command{
	Use:   "linkedlist",
	Short: "A linked list is a data structure of linearly linked elements",
	Long:  `This linked list contains pointers for head and tail`,
	Run: func(cmd *cobra.Command, args []string) {
		list := linkedlist.LinkedList(&linkedlist.SinglyLinkedList{})
		list.Prepend(linkedlist.NewNode(20))
		list.Prepend(linkedlist.NewNode(1))
		list.Prepend(linkedlist.NewNode(11))
		list.Prepend(linkedlist.NewNode(9))
		list.Prepend(linkedlist.NewNode(34))
		list.Print()
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
