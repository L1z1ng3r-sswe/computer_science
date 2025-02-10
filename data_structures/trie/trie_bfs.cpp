#include <iostream>
#include <unordered_map>

class Trie {
public:
  Trie() {
    m_children.reserve(26);
  }
  ~Trie () {
    for (auto& pair : m_children) {
      delete pair.second;
    }
  }

  void insert(std::string word) {
    Trie* currNode = this;
    
    for (const char& c : word ) {  
      if (currNode->m_children.find(c) == currNode->m_children.end()) {
        currNode->m_children[c] = new Trie();
      }
      currNode = currNode->m_children[c];
    }

    currNode->m_isWord = true;
  }

  bool search(std::string word) {
    Trie* currNode = this;
    
    for (const char& c : word ) {  
      if (currNode->m_children.find(c) == currNode->m_children.end()) {
        return false;
      }
      currNode = currNode->m_children[c];
    }

    return currNode->m_isWord;
  }

  bool findPrefix(std::string word) {
    Trie* currNode = this;
    
    for (const char& c : word ) {  
      if (currNode->m_children.find(c) == currNode->m_children.end()) {
        return false;
      }
      currNode = currNode->m_children[c];
    }

    return true;
  }

private:
  bool m_isWord = false;
  std::unordered_map<char, Trie*> m_children;
};