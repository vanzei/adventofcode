#ifndef UTILS_H
#define UTILS_H

#include <vector>
#include <string>

// Declaration of the function to read input from a file
std::vector<int> readInput(const std::string& filename);

// Declaration of the function to read commands from a file
std::vector<std::pair<std::string, int>> readCommands(const std::string& filename);

#endif // UTILS_H
