#include "utils.h"
#include <fstream>
#include <iostream>
#include <sstream>

// Definition of the function to read input from a file
std::vector<int> readInput(const std::string& filename) {
    std::vector<int> measurements;
    std::ifstream inputFile(filename);
    if (!inputFile) {
        std::cerr << "Error opening file: " << filename << std::endl;
        return measurements;
    }
    int value;
    while (inputFile >> value) {
        measurements.push_back(value);
    }
    return measurements;
}

// Definition of the function to read commands from a file
std::vector<std::pair<std::string, int>> readCommands(const std::string& filename) {
    std::vector<std::pair<std::string, int>> commands;
    std::ifstream inputFile(filename);
    if (!inputFile) {
        std::cerr << "Error opening file: " << filename << std::endl;
        return commands;
    }
    std::string line;
    while (std::getline(inputFile, line)) {
        std::istringstream iss(line);
        std::string direction;
        int value;
        if (!(iss >> direction >> value)) {
            continue; // Skip lines that don't match the expected format
        }
        commands.emplace_back(direction, value);
    }
    return commands;
}
