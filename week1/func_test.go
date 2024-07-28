package main

import "testing"

func TestDeferReturnV1(t *testing.T) {
	println(DeferReturnV1())
}
func TestDeferReturnV2(t *testing.T) {
	println(DeferReturnV2())
}
func TestDeferCloseLoopV1(t *testing.T) {
	DeferCloseLoopV1()
}

func TestDeferCloseLoopV2(t *testing.T) {
	DeferCloseLoopV2()
}

func TestDeferCloseLoopV3(t *testing.T) {
	DeferCloseLoopV3()
}
