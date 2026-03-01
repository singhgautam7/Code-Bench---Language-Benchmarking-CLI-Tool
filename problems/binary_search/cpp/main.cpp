#include <iostream>
#include <vector>
#include <sstream>
#include <string>

// Perform binary search on a sorted vector.
// Returns the index of target, or -1 if not found.
int binarySearch(const std::vector<int>& arr, int target) {
    int low = 0;
    int high = static_cast<int>(arr.size()) - 1;

    while (low <= high) {
        int mid = low + (high - low) / 2;
        if (arr[mid] == target) {
            return mid;
        } else if (arr[mid] < target) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return -1;
}

int main() {
    // Read all input from stdin
    std::string input;
    std::string line;
    while (std::getline(std::cin, line)) {
        input += line + " ";
    }

    std::istringstream iss(input);
    std::vector<int> tokens;
    int val;
    while (iss >> val) {
        tokens.push_back(val);
    }

    if (tokens.size() < 2) {
        std::cerr << "Error: expected array and target on stdin" << std::endl;
        return 1;
    }

    // Last token is target, rest are array
    int target = tokens.back();
    tokens.pop_back();

    int result = binarySearch(tokens, target);
    std::cout << result << std::endl;

    return 0;
}
