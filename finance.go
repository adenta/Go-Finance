package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func main() {

    file, e: = os.Open("./transactions.CSV")
    check(e)
    defer file.Close()

    path: = "./output.csv"
    var _, err = os.Stat(path)

    // if the output file is not found, create it.
    if os.IsNotExist(err) {
        var outFile, err = os.Create(path)
        check(err)
        defer outFile.Close()
    }
    outFile, err: = os.OpenFile(path, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
    check(err)



    labels: = make(map[string] string)
    userInput: = bufio.NewReader(os.Stdin)

    reader: = csv.NewReader(file)
    writer: = csv.NewWriter(outFile)
    reader.Comma = ','
    lineCount: = 0

    for {
        // read just one record, but we could ReadAll() as well
        record, err: = reader.Read()
            // end-of-file is fitted into err
        if err == io.EOF {
            break
        }
        check(err)

        fmt.Println(record)
        key: = record[4]
        price: = record[5]

        class: = make([] string, 3)

        if val, ok: = labels[key];
        ok {
            fmt.Print("already been classified as:" + val)
            class = strings.Split(strings.TrimSpace(labels[key]), ":")
            fmt.Print(class)

        } else {
            fmt.Print("Enter classification: ")
            text, _: = userInput.ReadString('\n')
            labels[key] = text

            class = strings.Split(strings.TrimSpace(text), ":")
        }

        row: = make([] string, 4)
        row[0] = key
        row[1] = price
        row[2] = class[0]

        if (len(class) > 1) {
            row[3] = class[1]

        }

        err = writer.Write(row)
        check(err)
        writer.Flush()

        lineCount += 1
    }
}
