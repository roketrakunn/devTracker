package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// StorageManager handles all file I/O operations
// Think of this as the "disk driver" - it knows HOW to read/write
// but doesn't care WHAT the data means
type StorageManager struct {
	baseDir string // Where we store all our files
}

// NewStorageManager creates a new storage manager
// By default, uses ~/.devtrack/ directory
func NewStorageManager() (*StorageManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	baseDir := filepath.Join(homeDir, ".devtrack") //create path
	
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}
	
	//return new NewStorageManager 
	return &StorageManager{
		baseDir: baseDir,
	}, nil
}

// NewStorageManagerWithPath creates a storage manager with a custom path
func NewStorageManagerWithPath(path string) (*StorageManager, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &StorageManager{
		baseDir: path,
	}, nil
}

// GetPath returns the full path for a given filename
func (sm *StorageManager) GetPath(filename string) string {
	return filepath.Join(sm.baseDir, filename)
}

// Read reads a JSON file and unmarshals it into the provided interface
// The 'v' parameter is a pointer to whatever struct you want to fill
func (sm *StorageManager) Read(filename string, v interface{}) error {
	path := sm.GetPath(filename)
	
	data, err := os.ReadFile(path)
	if err != nil {
		// If file doesn't exist, that's okay - return nil
		// (caller will get zero value of their struct)
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	// If file is empty, nothing to unmarshal
	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from %s: %w", filename, err)
	}

	return nil
}

// Write marshals the provided interface to JSON and writes it to a file
func (sm *StorageManager) Write(filename string, v interface{}) error {
	path := sm.GetPath(filename)

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// Exists checks if a file exists in storage
func (sm *StorageManager) Exists(filename string) bool {
	path := sm.GetPath(filename)
	_, err := os.Stat(path)
	return err == nil
}

// Delete removes a file from storage
func (sm *StorageManager) Delete(filename string) error {
	path := sm.GetPath(filename)
	
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file %s: %w", filename, err)
	}
	
	return nil
}

// BaseDir returns the base directory path
func (sm *StorageManager) BaseDir() string {
	return sm.baseDir
}
