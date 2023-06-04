package dsa

import (
	"fmt"
	"os"
	"path/filepath"
)

var DSAPath = filepath.FromSlash("src/DSA")

var DSAFiles = map[string][]string{
	"QuickSort":        {"QuickSort/QuickSort.go", "QuickSort/QuickSort_test.go"},
	"Queue":            {"Queue/Queue.go", "Queue/Queue_test.go"},
	"Stack":            {"Stack/Stack.go", "Stack/Stack_test.go"},
	"SinglyLinkedList": {"SinglyLinkedList/SinglyLinkedList.go", "SinglyLinkedList/SinglyLinkedList_test.go"},
	"DoublyLinkedList": {"DoublyLinkedList/DoublyLinkedList.go", "DoublyLinkedList/DoublyLinkedList_test.go"},
	"ArrayList":        {"ArrayList/ArrayList.go", "ArrayList/ArrayList_test.go"},
	"BubbleSort":       {"BubbleSort/BubbleSort.go", "BubbleSort/BubbleSort_test.go"},
	"LinearSearchList": {"LinearSearchList/LinearSearchList.go", "LinearSearchList/LinearSearchList_test.go"},
	"MergeSort":        {"MergeSort/MergeSort.go", "MergeSort/MergeSort_test.go"},
	"RingBuffer":       {"RingBuffer/RingBuffer.go", "RingBuffer/RingBuffer_test.go"},
	"MinHeap":          {"MinHeap/MinHeap.go", "MinHeap/MinHeap_test.go"},
	"TwoCrystalBalls":  {"TwoCrystalBalls/TwoCrystalBalls.go", "TwoCrystalBalls/TwoCrystalBalls_test.go"},
	"BinarySearchList": {"BinarySearchList/BinarySearchList.go", "BinarySearchList/BinarySearchList_test.go"},
	"MazeSolver":       {"MazeSolver/MazeSolver.go", "MazeSolver/MazeSolver_test.go"},
	"Trie":             {"Trie/Trie.go", "Trie/Trie_test.go"},
	"DFSOnBST":         {"DFSOnBST/DFSOnBST.go", "DFSOnBST/DFSOnBST_test.go"},
	"BTBFS":            {"BTBFS/BTBFS.go", "BTBFS/BTBFS_test.go"},
	"BTPreOrder":       {"BTPreOrder/BTPreOrder.go", "BTPreOrder/BTPreOrder_test.go"},
}

func Copy(dsa string, dstDir string) error {
	dsaFiles, ok := DSAFiles[dsa]
	if !ok {
		return fmt.Errorf("algorithm %s could not be found", dsa)
	}

	for _, dsaFile := range dsaFiles {
		file, err := os.ReadFile(filepath.Join(DSAPath, dsaFile))
		if err != nil {
			return fmt.Errorf("read file %s: %w", dsaFile, err)
		}

		dst := filepath.Join(dstDir, filepath.FromSlash(dsaFile))
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			return fmt.Errorf("make dirs for %s: %w", dst, err)
		}
		if err := os.WriteFile(dst, file, 0644); err != nil {
			return fmt.Errorf("write file %s: %w", dst, err)
		}
	}

	return nil
}
