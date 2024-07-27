#include <iostream>
#include <vector>
#include <string>
#include "utils.h"

// Function to process commands and calculate final horizontal position and depth
std::pair<int, int> processCommands(const std::vector<std::pair<std::string, int>>& commands) {
    int horizontalPosition = 0;
    int depth = 0;
    for (const auto& command : commands) {
        const std::string& direction = command.first;
        int value = command.second;
        if (direction == "forward") {
            horizontalPosition += value;
        } else if (direction == "down") {
            depth += value;
        } else if (direction == "up") {
            depth -= value;
        }
    }
    return {horizontalPosition, depth };
}

int main(int argc, char* argv[]) {
    if (argc < 2) {
        std::cerr << "Usage: " << argv[0] << " <input file path>" << std::endl;
        return 1;
    }
    std::string filename = argv[1];
    auto commands = readCommands(filename);
    if (commands.empty()) {
        std::cerr << "No valid commands found in the input file." << std::endl;
        return 1;
    }
    auto [horizontalPosition, depth] = processCommands(commands);
    std::cout << "Final Horizontal Position: " << horizontalPosition << std::endl;
    std::cout << "Final Depth: " << depth << std::endl;
    std::cout << "Result: " <<  horizontalPosition * depth << std::endl;
    return 0;
}
