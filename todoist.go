package main

import (
  "os"
	"github.com/jawher/mow.cli"
)

/* *************************************** */
/* MAIN ********************************** */
/* *************************************** */
func main() {
  app := cli.App("todoist", "Command line client for Todoist.com, ease-to-use tasks manager")


/*  command := arguments[0]
  _, err := getSavedToken()

  if command != "init" && err != nil {
    ct.ChangeColor(ct.Red, false, ct.None, false)
    fmt.Println("Todoist program doesn't have a valid user token. You must launch `todoist init [token]` to save your authentication token.")
    ct.ResetColor()
    os.Exit(1)
  }*/


  app.Command("init", "initialize the todoist program to be connected to your todoist.com account", func(cmd *cli.Cmd) {
    cmd.Spec = "[TOKEN]"
    token := cmd.StringArg("TOKEN", "", "Your personal API token that your can find into your Todoist account: Settings -> Todoist Settings -> Account -> API token")
    cmd.Action = func() {
      askUserForToken(*token)
    }
  })

  app.Command("list", "list some data (projects, tasks, etc.) from your Todoist account", func(cmd *cli.Cmd) {
    cmd.Command("projects", "list all your projects", func(cmd *cli.Cmd) {
      cmd.Action = listProjects
    })
    cmd.Command("tasks", "list all your tasks", func(cmd *cli.Cmd) {
      cmd.IntOpt("p project", 0, "The project you want to retrieve tasks. You must only write the number corresponding with the project. You can find all projects with their number executing `todoist list projects`")
      cmd.Action = listTasks
    })
    cmd.Command("labels", "list all your labels", func(cmd *cli.Cmd) {
      cmd.Action = listLabels
    })
  })

  app.Command("add", "add a task", func(cmd *cli.Cmd) {
    projectOrder := cmd.IntOpt("p project", 0, "The project you want to add the task. You must only write the number corresponding with the project. You can find all projects with their number executing `todoist list projects`")
    priority := cmd.IntOpt("priority", 1, "The priority of the task (a number between 1 and 4, 4 for very urgent and 1 for natural).")
    date := cmd.StringOpt("d date", "", "The date of the task, added in free form text, for example it can be every day @ 10. Look at our reference to see which formats are supported. We are currently assuming that dates are written in English. https://todoist.com/Help/DatesTimes")
    labels := cmd.StringOpt("l labels", "", "The task labels")
    task := cmd.StringArg("TASK", "", "The task you want to add.")
    cmd.Action = func() {
      addTask(*task, *projectOrder, *date, *labels, *priority)
    }
  })

  app.Command("task", "process a task", func(cmd *cli.Cmd) {
    task := cmd.IntArg("TASK_ID", 0, "The identifier of the task you want to process.")
    cmd.Command("done", "Mark the task as done", func(cmd *cli.Cmd) {
      cmd.Action = func() {
        markTaskAsDone(*task)
      }
    })
    cmd.Command("undone", "Mark the task as undone", func(cmd *cli.Cmd) {
      cmd.Action = func() {
        markTaskAsUndone(*task)
      }
    })
  })

  app.Run(os.Args)
}
