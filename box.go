package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errOutOfRange  = errors.New("out of the shapesCapacity range")
	errInxNotExist = errors.New("index does not exist")
	errEmptyList   = errors.New("empty list")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return fmt.Errorf("%w", errOutOfRange)
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	for j, shape := range b.shapes {
		if j == i {
			return shape, nil
		}
	}
	return nil, errInxNotExist

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) < i || len(b.shapes) == i {
		return nil, errInxNotExist
	}

	removed := b.shapes[i]
	if len(b.shapes)-1 == i {
		b.shapes = b.shapes[:i]
		return nil, nil
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return removed, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var removed Shape
	if len(b.shapes) <= i {
		return nil, errInxNotExist
	} else if len(b.shapes)-1 >= i {
		removed = b.shapes[i]
		b.shapes[i] = shape
	}

	return removed, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var ss float64
	for _, shape := range b.shapes {
		ss += shape.CalcPerimeter()
	}
	return ss

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, shape := range b.shapes {
		sumArea += shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	if len(b.shapes) == 0 {
		return errEmptyList
	}
	var newShapes []Shape
	for _, shape := range b.shapes {
		if _, ok := shape.(*Circle); !ok {
			newShapes = append(newShapes, shape)
		}
		if len(newShapes) == len(b.shapes) {
			return errEmptyList
		}
	}

	b.shapes = newShapes

	return nil

}
