# search
Search for a file/folder/both in go

Pretty basic usage:

```
package main

func main () {
	directories := search.Search("filename", "filepath", "folder") // replace folder with 'file' or 'all' if you wish

	for _, element := range directories {
		os.RemoveAll(element)   // deletes everything it found
	}
}

```
