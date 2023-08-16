package main

import "fmt"

type Shape interface {
	getShape()
}
type Circle struct {
}

func (r *Circle) getShape() {
	fmt.Println("I am a circle")
}

type Square struct {
}

func (s *Square) getShape() {
	fmt.Println("I am a square")
}

type Color interface {
	getColor()
}
type Red struct {
}

func (r *Red) getColor() {
	fmt.Println("I am red")
}

type Blue struct {
}

func (b *Blue) getColor() {
	fmt.Println("I am blue")
}

type AbstractFactory interface {
	getShape(shape string) *ShapeFactory
	getColor(color string) *ColorFactory
}

type ShapeFactory struct {
}

func (s *ShapeFactory) getShape(shape string) *ShapeFactory {
	if shape == "circle" {
		return (*ShapeFactory)(&Circle{})
	}
	if shape == "square" {
		return (*ShapeFactory)(&Square{})
	}
	return nil
}
func (s *ShapeFactory) getColor(color string) *ColorFactory {
	return nil
}

type ColorFactory struct {
}

func (c *ColorFactory) getColor(color string) *ColorFactory {
	if color == "red" {
		return (*ColorFactory)(&Red{})
	}
	if color == "blue" {
		return (*ColorFactory)(&Blue{})
	}
	return nil
}
func (c *ColorFactory) getShape(shape string) *ShapeFactory {
	return nil
}

//工厂创建器
type FactoryProducer struct {
}

func (f *FactoryProducer) getShapeFactory(shape string) *ShapeFactory {
	if shape == "CIRCLE" {
		return (*ShapeFactory)(&Circle{})
	}
	if shape == "SQUARE" {
		return (*ShapeFactory)(&Square{})
	}
	return nil
}
func (f *FactoryProducer) getColorFactory(color string) *ColorFactory {
	if color == "RED" {
		return (*ColorFactory)(&Red{})
	}
	if color == "BLUE" {
		return (*ColorFactory)(&Blue{})
	}
	return nil
}
func main() {

	//factory := &FactoryProducer{}.getShapeFactory("CIRCLE")
}
