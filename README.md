### Compression Tool 🗜️

A compression tool is a piece of software that is used to store a file data in a
compressed format. It is useful for space optimization i.e it helps to store the
same amount of information using less memory space. In this application Huffman
coding algorithm is used to compress the file content.

### Huffman Encoding Algorithm 🚀

The main steps of the algorithm are explained below -

1. **Frequency Count**: First and most important step of this algorithm is to
   count the frequency of each character present in the source file.
2. **Build Huffman Tree**: Next step is to build a Huffman tree based on the
   frequency table.
3. **Encode Tree**: We assign the left edge a value of `0` and right `1` for
   each node from root node till we reach a child node. After this process we
   can obtain the encoded bit string for each character. Since characters are
   present in the child node only, the route from root to a particular child
   node is the encoded bit string for that particular character.

**Let's understand the Algorithm with the help of an example** 🧠

1. Suppose the file content is `halamadrid`
2. The frequency table would look like this.

    | Character | Frequency |
    | -------------- | --------------- |
    | `h` | 1 |
    | `a` | 3 |
    | `l` | 1 |
    | `m` | 1 |
    | `d` | 2 |
    | `r` | 1 |
    | `i` | 1 |
3. The Huffman Tree (🌲) for the corresponding table would be 

<img src="./huffman_tree.png" height=400 />

4. After building the tree, the bit encoding of a Character is equal to the path from root node to a child node. The encodings are shown below - 

    | Character | Encoding |
    | -------------- | --------------- |
    | `h` | 000 |
    | `a` | 11 |
    | `l` | 001 |
    | `m` | 1000 |
    | `d` | 01 |
    | `r` | 1001 |
    | `i` | 101 |

5. So, the encoded bit sequence for `halamadrid` will be `000110011110001101100110101`

### Application Specific Details 
- First build the application by running `make`.
- Run the application using `./gc` or by running `go run main.go` directly.

### Output
<img src="./gcomp_output.png" />
<br/>
<br/>

In the above picture, it can be seen that the compressed file `gc_comp.txt` occupies less memory than the source file `gc_source.txt`.
