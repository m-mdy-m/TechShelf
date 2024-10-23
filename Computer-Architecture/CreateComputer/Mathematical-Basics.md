### **Mathematical Basics for Computer Architecture (8-bit Computer)**

#### 1. **[Binary Number System](https://newcollege.ac.in/CMS/Eknowledge/73af25af-c58f-49b2-9621-a1e07d01ac7bNumber%20System.pdf)**

- **Definition**: Computers use binary (base-2) number system. All operations, data storage, and communication in computers are done in binary.
- **Concepts**:
  - **Bit**: A single binary digit (0 or 1).
  - **Byte**: A group of 8 bits, the standard unit of data in 8-bit computers.
- **Conversion**:
  - Decimal to binary and vice versa (important for understanding data representation).
- **Binary Arithmetic**: Addition, subtraction, multiplication, and division in binary.

#### 2. **[Boolean Algebra](https://www.hzu.edu.in/csit/Boolean%20Algebra%20%20computer%20fundamentals.pdf)**

- **Definition**: A branch of algebra used to work with binary numbers and logical operations.
- **Basic Boolean Operations**:
  - **AND**: Output is 1 if both inputs are 1.
  - **OR**: Output is 1 if at least one input is 1.
  - **NOT**: Inverts the input (0 becomes 1, 1 becomes 0).
- **Combinational Circuits**: Using Boolean algebra to design logic circuits such as gates (AND, OR, NOT, XOR).
- **De Morgan’s Laws**: Important for simplifying logical expressions.
- [Boolean algebra and Logic Gates](https://www.pvpsiddhartha.ac.in/dep_it/lecture%20notes/DSD/unit2.pdf)

#### 3. **[Logic Gates](https://epgp.inflibnet.ac.in/epgpdata/uploads/epgp_content/S000574EE/P001494/M015065/ET/1459848930et05.pdf) [and Circuits](https://www.cs.cornell.edu/courses/cs3410/2019sp/schedule/slides/02-gates-notes.pdf)**

- **Logic Gates**: Physical implementations of Boolean operations.
  - **Basic Gates**: AND, OR, NOT, NAND, NOR, XOR, XNOR.
- **Truth Tables**: Representation of the output of a logic gate for all input combinations.
- **Combining Gates**: Building larger circuits by combining simple gates to perform complex operations.
  - **Examples**: Half adder, full adder (used for binary addition in ALUs).

#### 4. **[Arithmetic Logic Unit (ALU) Design](https://www.csie.ntu.edu.tw/~cyy/courses/introCS/17fall/lectures/handouts/lec04_ALU.pdf)**

- **Binary Addition**: Using half and full adders to design the ALU for basic arithmetic.
- **Two’s Complement Arithmetic**: For representing negative numbers and performing subtraction.
- **Bitwise Operations**: AND, OR, XOR, NOT (used in ALU for logical operations).
- **Shift Operations**: Left shift and right shift, used for multiplication and division by powers of 2.
- [https://epgp.inflibnet.ac.in/epgpdata/uploads/epgp_content/S000574EE/P001494/M019814/ET/1491915021p4m15_etext.pdf]
- [https://faculty.ksu.edu.sa/sites/default/files/Unit-10%20ALU%20Design.pdf]
- [https://www.rvstcc.ac.in/assets/img/pdf/unit-1.pdf]
- [https://courses.cs.duke.edu/spring05/cps104/lectures/2up-lecture09.pdf]
- [https://safari.ethz.ch/digitaltechnik/spring2023/lib/exe/fetch.php?media=lab5_manual.pdf]
- [https://www.cs.uic.edu/~i266/fall12_hw10/5571.pdf]

#### 5. **[Memory Addressing and Binary Representation](https://courses.cs.washington.edu/courses/cse351/17wi/lectures/CSE351-L02-memory-I_17wi.pdf)**

- [Memory Addressing](https://www.analog.com/media/en/training-seminars/design-handbooks/microprocessor-systems-handbook/Chapter2.pdf)
- **Memory**: Organized as an array of bytes (in an 8-bit computer, each memory location holds 8 bits).
- **[Address Bus](https://cds.cern.ch/record/872217/files/p156.pdf)**: The number of address lines determines the number of memory locations (2^n memory locations for n address lines).
  - **Example**: 8-bit address bus can access 2^8 = 256 memory locations.
- **Data Bus**: Transmits data between memory and CPU.
  - **Example**: 8-bit data bus can transfer 1 byte at a time.
- **Memory Mapping**: Assigning specific addresses to different devices (ROM, RAM, I/O).

#### 6. **Clock Cycles and Timing Diagrams**

- **Clock Signal**: Synchronizes the operations of the CPU. Each operation takes a certain number of clock cycles.
- **[Timing Diagrams](https://www.skdavpolytech.ac.in/news_files/microprocessor_part_2_compressed_1588258956.pdf)**: Graphical representations of the relationship between the clock signal and various signals (control, data, etc.) in the CPU.
- **Calculation**: Understanding the clock frequency and calculating the time it takes to execute instructions (Clock Speed = 1 / Time Period).

#### 7. **Finite State Machines (FSMs)**

- **Definition**: A computational model used to design control units in CPUs.
- **States**: A finite set of states, transitions between them based on inputs (used in control logic).
- **Mathematical Representation**:
  - **State Transition Table**: Describes how the state changes in response to inputs.
  - **State Diagrams**: Visual representation of states and transitions.

#### 8. **Binary Multiplication and Division**

- **Binary Multiplication**: Similar to decimal multiplication, involves bitwise shifts and additions.
- **Binary Division**: Repeated subtraction method or shift-and-subtract method.

#### 9. **Hexadecimal and Octal Systems**

- **Hexadecimal (Base-16)**: Used for compact representation of binary numbers.
  - **Conversions**: Binary to hexadecimal, hexadecimal to binary.
  - **Example**: 1 byte = 8 bits = 2 hexadecimal digits.
- **Octal (Base-8)**: Another shorthand for binary representation (grouping bits in sets of 3).

#### 10. **Instruction Set Architecture (ISA) and Opcodes**

- **Instruction Format**: 8-bit computers have instructions encoded in binary, called **opcodes**.
  - Example: MOV, ADD, SUB (instructions with their binary representation).
- **Instruction Cycle**:
  - **Fetch**: The CPU fetches the instruction from memory.
  - **Decode**: The CPU decodes the binary instruction into operations.
  - **Execute**: The CPU performs the operation.
- **Instruction Set Design**: Designing binary codes for each operation (typically involves some combinatorics to assign unique codes).

#### 11. **Addressing Modes**

- **Direct Addressing**: The operand is directly specified in the instruction.
- **Indirect Addressing**: The address of the operand is stored in a memory location.
- **Indexed Addressing**: A base address is added to an offset to get the operand’s location.
- **Immediate Addressing**: The operand is part of the instruction itself.
- **Mathematical Calculation**: Understanding how the effective address is calculated in different modes.

#### 12. **Bus Architecture and Data Transfer**

- **Bus**: A set of parallel wires that carry data, address, and control signals.
- **Control Signals**: Read/Write, clock, reset, and interrupt signals.
- **Tri-State Buffers**: Allow multiple devices to share a bus.
- **Bus Timing Diagrams**: Analyzing how data is transferred in relation to clock cycles and control signals.

#### 13. **Error Detection and Correction**

- **Parity Bit**: A simple error detection method where a bit is added to data to make the number of 1’s either odd or even.
- **Hamming Code**: An error-correcting code that detects and corrects single-bit errors.

#### 14. **CPU Control Unit Design**

- **Control Signals**: Generated by the control unit to direct the operation of other parts of the CPU.
- **Microinstructions**: Low-level instructions that define each step of the instruction execution.
- **Design of Control Logic**: Using state machines to generate control signals based on the current instruction and the state of the CPU.

#### 15. **Power of Two and Memory Sizes**

- **Calculations**: Understanding powers of two to calculate memory sizes and limits.
  - Example: 2^8 = 256 (maximum number of unique addresses in an 8-bit address space).
- **Kilobytes, Megabytes, and Gigabytes**: Understanding memory sizes in powers of 1024.
