# **CS(computer science)-CreateComputer**

## What is CreateComputer?

Basics and prerequisites and how to build an 8-bit computer

## Recommended Books and Resources

### Books

### **Foundational Knowledge (Prerequisites)**

#### 1. **Digital Logic Design**

- **Understanding the basic building blocks** of CPUs, such as logic gates, flip-flops, counters, and registers is essential for designing an 8-bit computer.
- **Book**: [Digital Design and Computer Architecture](https://github.com/kaitoukito/Computer-Science-Textbooks/blob/master/Digital-Design-and-Computer-Architecture-RISC-V-Edition.pdf) by David Harris and Sarah Harris. (Highly recommended for understanding how logic gates and circuits form the basis of a CPU.)

#### 2. **CPU Architecture and Design**

- Learn about the architecture of CPUs: how registers, ALUs, and control units work together.
- **Book**: [Computer Organization and Design: The Hardware/Software Interface](https://theswissbay.ch/pdf/Books/Computer%20science/Computer%20Organization%20and%20Design-%20The%20HW_SW%20Inteface%205th%20edition%20-%20David%20A.%20Patterson%20%26%20John%20L.%20Hennessy.pdf) by David A. Patterson and John L. Hennessy. (Covers basic CPU architecture in depth.)

#### 3. **Electronics Fundamentals**

- You need to understand basic electronics principles (Ohm’s law, voltage, current) to work with physical components like resistors, capacitors, and transistors.
- **Book**: [The Art of Electronics](https://eclass.uniwa.gr/modules/document/index.php?course=EEE265&download=/6218412aGLbn/5e6db8cd5KNP.pdf) by Paul Horowitz and Winfield Hill. (The go-to book for electronics engineering.)

#### 4. **Assembly Language Programming**

- Since you'll program the CPU in assembly, it's crucial to understand how assembly code works on an 8-bit architecture.
- **Book**: [Programming the Z80](http://www.z80.info/zip/programming_the_z80_3rd_edition.pdf) by Rodnay Zaks. (A definitive guide for Z80 assembly.)
- [Programming the Z80", 3rd Edition](https://seriouscomputerist.atariverse.com/media/pdf/book/Z80%20Assembly%20Language%20Programming.pdf)

### **Detailed Resources for Building an 8-bit PC**

#### 1. **Choosing the CPU or Designing Your Own**

- **CPU Options**: You could choose to use a classic 8-bit CPU like the Z80 or 6502, or design your own using logic gates.
- **Resource**: [Ben Eater’s 8-bit Breadboard Computer Series](https://eater.net/8bit) on YouTube (A fantastic step-by-step guide on building an 8-bit CPU from scratch using breadboards, including the clock, ALU, and control logic).
- **Forum Thread**: [6502.org](http://www.6502.org/) is a community focused on building computers using the 6502 processor and includes design tutorials, datasheets, and programming tips.
- **Guide**: [Build an 8-bit CPU from Scratch](https://theoldnet.com/museum/build_8_bit_cpu/) (detailed guide that explains how to design and implement your own 8-bit CPU with diagrams).

#### 2. **Designing and Implementing the Bus and Memory Architecture**

- **Memory Tutorial**: [Memory Mapping in an 8-bit System](https://projects.drogon.net/6502/memory-mapping/) by drogon.net.
- **Book**: [The Z80 Microprocessor: Architecture, Interfacing, Programming and Design](http://www.primrosebank.net/computers/mtx/projects/mtxplus/data/books/TheZ80Microprocessor.pdf) by Ramesh Gaonkar (Covers the Z80 CPU's memory and I/O interfacing techniques).
- **Memory Interfacing Guide**: [Designing a Memory System](https://www.tayloredge.com/reference/Electronics/memorydesign.pdf) (Practical guide for interfacing SRAM and EEPROM with an 8-bit processor).

#### 3. **Building and Designing the Circuit (Breadboard or PCB)**

- **Guide**: [Wiring and Building an 8-bit CPU on Breadboards](https://projects.nikolasdesign.com/microcomputer-on-a-breadboard/) (A practical walkthrough to wiring all components including clock, memory, and buses on a breadboard).
- **Breadboard Tutorials**: [Breadboard Basics](https://www.evilmadscientist.com/2008/beginning-breadboarding/) by Evil Mad Scientist (Covers practical tips for working with breadboards).
- **Tool**: [KiCad](https://kicad.org/) (Open-source tool for designing PCB layouts if you want to move from breadboarding to PCB manufacturing).

#### 4. **Programming the BIOS, Bootloader, and Writing Assembly Programs**

- **Resource**: [Writing a Simple BIOS for 8-bit Computers](https://www.edn.com/biographies-for-the-8080-and-z80-processors/) (Explains how to write a simple BIOS that initializes hardware and loads the first program).
- **Assembler**: [TASM Assembler](https://www.emu8086.com/tasm/) (A simple assembler to convert assembly code into machine code for your 8-bit computer).
- **Tool**: [SimH](https://github.com/simh/simh) (An emulator that can run 8-bit CPU architectures, useful for debugging before you load programs onto hardware).

#### 5. **Building Input/Output (I/O) Systems**

- **Guide**: [Implementing a Simple I/O Interface](https://cs.uwaterloo.ca/~brecht/servers/docs/PowerEdge-2600/en/Pe6450/UG/3097Pab0.pdf) (How to create input and output systems like LEDs or serial terminals).
- **Project**: [VGA Output from an 8-bit CPU](https://eater.net/vga) (A tutorial by Ben Eater to add VGA video output to your custom 8-bit computer).
- **Book**: [Interfacing to the Z80](https://www.zilog.com/docs/z80/um0080.pdf) (Comprehensive coverage of interfacing peripherals to the Z80 and similar CPUs).

### **Hands-on Projects and Examples**

#### 1. **Ben Eater’s 8-bit Computer Series**

- **YouTube Series**: [Building an 8-bit CPU from Scratch](https://www.youtube.com/watch?v=HyznrdDSSGM) (The go-to video tutorial series that explains how to build a working 8-bit computer on a breadboard, covering everything from the CPU to memory and I/O.)

#### 2. **Making a Retro 8-bit Computer with the Z80**

- **Blog**: [Z80 Computer Build Blog](https://pdfroom.com/books/build-your-own-z80-computer-steve-ciarcia/X6234b3654Z/download) (Step-by-step guide to building a Z80-based 8-bit computer, including schematic diagrams and programming the ROM).
- **Forum**: [RetroComputing StackExchange](https://retrocomputing.stackexchange.com/questions/4825/how-to-build-an-8-bit-computer) (A community with discussions on building 8-bit computers, solving problems, and sharing knowledge).

---

### **Additional**
