package functions

import "strings"

func removeComments(line string, inBlockComment *bool) string {
    var result strings.Builder

    for i := 0; i < len(line); {
        if *inBlockComment {
            end := strings.Index(line[i:], "*/")

            if end == -1 {
                return result.String()
            }

            i += end + 2
            *inBlockComment = false
            continue
        }

        if i+1 < len(line) && line[i:i+2] == "//" {
            break
        }

        if i+1 < len(line) && line[i:i+2] == "/*" {
            *inBlockComment = true
            i += 2
            continue
        }

        result.WriteByte(line[i])
        i++
    }

    return result.String()
}