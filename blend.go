package blender

import "errors"

func Blend(substance string, inputLiter float64) (output string, outputLiter float64, err error) {
	switch substance {
	case "apple":
		{
			output = "apple juice"
			outputLiter = inputLiter * 0.8
		}
	case "water melon":
		{
			output = "water melon juice"
			outputLiter = inputLiter * 0.8
		}
	case "celery":
		{
			output = "celery juice"
			outputLiter = inputLiter * 0.8
		}
	case "water":
		{
			output = "water"
			outputLiter = inputLiter
		}
	case "apple juice":
		{
			output = "apple juice"
			outputLiter = inputLiter
		}
	case "water melon juice":
		{
			output = "water melon juice"
			outputLiter = inputLiter
		}
	case "celery juice":
		{
			output = "celery juice"
			outputLiter = inputLiter
		}
	case "rock":
		{
			err = errors.New("broken blender")
		}
	case "panda":
		{
			err = errors.New("wtf")
		}
	default:
		err = errors.New("unrecognized substance")
	}
	return output, outputLiter, err
}
