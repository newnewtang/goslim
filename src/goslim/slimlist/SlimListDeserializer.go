package slimlist

import (
	"errors"
	"regexp"
	"strconv"
)

func skipByte(current []byte, idx *int, expected byte) error {
	if current == nil || *idx >= len(current) || current[*idx] != expected {
		return errors.New("unexpected charactor!")
	}
	*idx += 1
	return nil
}

func readLength(current []byte, idx *int) int {
	length, _ := strconv.Atoi(string(current[*idx : *idx+6]))
	*idx += 6
	return length
}

func getByteLength(characterLength int, current []byte, idx int) int {
	chars := 0
	bytes := 0
	for p := idx; chars <= characterLength; p++ {
		bytes++
		if CSlim_IsCharacter(current[p]) {
			chars++
		}
	}
	if chars > characterLength {
		bytes--
	}
	return bytes
}

func SlimList_Deserialize(serializedList string) *SlimList {
	if len(serializedList) < 7 {
		return nil
	}

	current := []byte(serializedList)
	idx := 0
	list := SlimList_Create()

	if skipByte(current, &idx, '[') != nil {
		return nil
	}

	listLength := readLength(current, &idx)

	if skipByte(current, &idx, ':') != nil {
		return nil
	}

	for k := 0; k < listLength; k++ {
		characterLength := readLength(current, &idx)
		if skipByte(current, &idx, ':') != nil {
			return nil
		}

		byteLength := getByteLength(characterLength, current, idx)
		list.AddString(string(current[idx : idx+byteLength]))

		idx += byteLength

		if skipByte(current, &idx, ':') != nil {
			return nil
		}
	}

	if skipByte(current, &idx, ']') != nil {
		return nil
	}
	return list
}

func parseHashEntry(row string) *SlimList {
	cellsRegex := regexp.MustCompile(`<td>(.*?)</td>`)
	cells := cellsRegex.FindAllStringSubmatch(row, -1)

	if cells != nil && len(cells) == 2 && len(cells[0]) == 2 && len(cells[1]) == 2 {
		element := SlimList_Create()
		hashKey := cells[0][1]
		hashValue := cells[1][1]
		element.AddString(hashKey)
		element.AddString(hashValue)
		return element
	}
	return nil
}

func SlimList_deserializeHash(serializedHash string) *SlimList {
	hash := SlimList_Create()

	rowsRegex := regexp.MustCompile(`<tr>.*?</tr>`)

	rows := rowsRegex.FindAllString(serializedHash, -1)
	for _, v := range rows {
		hash.AddList(parseHashEntry(v))
	}

	return hash
}
