// @Description nestedCode
// @Author caopengfei 2022/2/7 15:16
// if嵌套代码，是新人难以理解代码的问题，实际编程中需要减少if的嵌套
// 在if 中直接return 减少 else的出现

package nestedCode

import "errors"

func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil

}

func concatenate(s1 string, s2 string) (s string, err error) {
	// ...
	return
}
