package main

import (
        "bufio"
        "encoding/csv"
        "encoding/json"
        "encoding/xml"
        "fmt"
        "log"
        "net/http"
        "os"
        "path/filepath"
        "strconv"
        "strings"
)

func main() {
        // ポート番号の入力を求める
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("port number: ")
        portStr, _ := reader.ReadString('\n')
        portStr = strings.TrimSpace(portStr)
        port, err := strconv.Atoi(portStr)
        if err != nil {
                log.Fatalf("ポート番号の変換に失敗しました: %v", err)
        }

        // ファイル名の入力を求める
        fmt.Print("file name: ")
        fileName, _ := reader.ReadString('\n')
        fileName = strings.TrimSpace(fileName)

        // 現在のディレクトリを取得
        currentDir, err := os.Getwd()
        if err != nil {
                log.Fatalf("現在のディレクトリの取得に失敗しました: %v", err)
        }

        // ハンドラ関数
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                filePath := filepath.Join(currentDir, fileName)
                ext := filepath.Ext(fileName)
                switch ext {
                case ".txt":
                        // テキストファイル
                        http.ServeFile(w, r, filePath)
                case ".csv":
                        // CSVファイル
                        serveCSV(w, filePath)
                case ".json":
                        // JSONファイル
                        serveJSON(w, filePath)
                case ".xml":
                        // XMLファイル
                        serveXML(w, filePath)
                default:
                        // その他のファイル（HTMLなど）
                        http.ServeFile(w, r, filePath)
                }
        })

        // サーバー起動
        addr := fmt.Sprintf(":%d", port)
        fmt.Printf("サーバーをポート%dで起動します...\n", port)
        log.Fatal(http.ListenAndServe(addr, nil))
}

func serveCSV(w http.ResponseWriter, filePath string) {
        file, err := os.Open(filePath)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        defer file.Close()

        reader := csv.NewReader(file)
        records, err := reader.ReadAll()
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "text/csv")
        for _, record := range records {
                fmt.Fprintf(w, strings.Join(record, ",")+"\n")
        }
}

func serveJSON(w http.ResponseWriter, filePath string) {
        file, err := os.Open(filePath)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        defer file.Close()

        decoder := json.NewDecoder(file)
        var data interface{}
        err = decoder.Decode(&data)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        encoder := json.NewEncoder(w)
        encoder.SetIndent("", "  ")
        err = encoder.Encode(data)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
}

func serveXML(w http.ResponseWriter, filePath string) {
        file, err := os.Open(filePath)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        defer file.Close()

        decoder := xml.NewDecoder(file)
        var data interface{}
        err = decoder.Decode(&data)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "application/xml")
        encoder := xml.NewEncoder(w)
        encoder.Indent("", "  ")
        err = encoder.Encode(data)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
}