package todo

import (
	"encoding/json"
	"path/filepath"
	// "path/filepath"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Stringer interface {
	String() string
}

type List []item

func (l *List) Add(task string) {
	i := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, i)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
  exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
  basepath := filepath.Dir(exePath)
  if err != nil {
    log.Println(err)
  }

  dataDir := filepath.Join(basepath, "data")
  if _, err := os.Stat(dataDir); os.IsNotExist(err){
    err := os.Mkdir(dataDir, os.ModePerm) 
      if err != nil{
      log.Println("Error creando el directorio 'data':", err)
			return err
    }
  }
  // Guardar el archivo dentro de "data/"
	filePath := filepath.Join(dataDir, filename)

	// Escribir el archivo JSON
	err = os.WriteFile(filePath, js, 0644)
	if err != nil {
		log.Println("Error escribiendo el archivo JSON:", err)
		return err
	}

	fmt.Println("Archivo guardado en:", filePath)
	return nil
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "[ ]"
		if t.Done {
			prefix = "[X] "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%d.%s: %s\n", k+1, prefix, t.Task)
	}
	return formatted
}

func (l *List) StringTime() string {
	formatted := ""
	for k, t := range *l {
		prefix := "[ ]"
		if t.Done {
			prefix = "[X] "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s, %d.%s: %s\n", t.CreatedAt.Format("02-01-2006 15:04:05"), k+1, prefix, t.Task)
	}
	return formatted
}

func (l *List) Pending() string {
	formatted := ""
	for k, t := range *l {
		prefix := "[ ]"
		if t.Done {
			continue
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s, %d.%s: %s\n", t.CreatedAt.Format("02-01-2006 15:04:05"), k+1, prefix, t.Task)
	}
	return formatted
}
