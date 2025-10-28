package main

import "fmt"

func reverseWords(s string) string {
    var words []string
    word := ""

    // Разделяем строку вручную
    for i := 0; i < len(s); i++ {
        ch := s[i]
        if ch != ' ' {
            word += string(ch)
        } else if word != "" {
            // добавляем слово и очищаем буфер
            words = append(words, word)
            word = ""
        }
    }

    // Добавляем последнее слово, если есть
    if word != "" {
        words = append(words, word)
    }

    // Склеиваем в обратном порядке вручную
    result := ""
    for i := len(words) - 1; i >= 0; i-- {
        result += words[i]
        if i != 0 {
            result += " "
        }
    }

    return result
}

func reverseWordsWithLib(s string) string{
    bankWord := strings.Split(s, " ")
    var reversed []string
    for i := len(bankWord)-1; i >= 0; i-- {
        if bankWord[i] == "" {
            continue
        }
        reversed = append(reversed, bankWord[i])
    }
    return strings.Join(reversed, " ")
}


func main() {
    fmt.Println(reverseWords("  hello   world  go  ")) // "go world hello"
}
