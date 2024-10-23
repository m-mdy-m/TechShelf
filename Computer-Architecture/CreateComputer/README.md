# **Computer-Architecture-CreateComputer**

## What is CreateComputer?

"CreateComputer" refers to the process of designing and building a basic, functioning computer from the ground up, typically an 8-bit computer, for educational or experimental purposes. This involves understanding how a computer works at the most fundamental level—starting with digital logic design, continuing through CPU architecture, memory management, input/output systems, and assembly language programming.

This guide will help you through the steps required to build an 8-bit computer, with a focus on essential concepts like logic gates, CPU architecture, memory, and bus design. It provides resources to acquire foundational knowledge, followed by practical guidance on building your own computer.

## Recommended Books and Resources

### Books

## Recommended Books and Resources

### **Foundational Knowledge (Prerequisites)**

Before you begin creating an 8-bit computer, it is important to acquire some foundational knowledge.

#### 1. **Digital Logic Design**

- **Understanding the basic building blocks**: Logic gates, flip-flops, counters, and registers form the foundation of any CPU. Grasping how these elements combine to create more complex digital circuits is essential.

  - **Book**: [**Digital Design and Computer Architecture**](https://github.com/kaitoukito/Computer-Science-Textbooks/blob/master/Digital-Design-and-Computer-Architecture-RISC-V-Edition.pdf) by David Harris and Sarah Harris  
    This book provides an excellent introduction to digital circuits and architecture, offering a clear path from understanding logic gates to building a functioning computer system. It’s highly recommended for anyone wanting to understand how logic circuits form the backbone of CPU architecture.

#### 2. **CPU Architecture and Design**

- **CPU components**: Learn how key components like registers, the Arithmetic Logic Unit (ALU), and control units work together to execute instructions.

  - **Book**: [**Computer Organization and Design: The Hardware/Software Interface**](https://theswissbay.ch/pdf/Books/Computer%20science/Computer%20Organization%20and%20Design-%20The%20HW_SW%20Inteface%205th%20edition%20-%20David%20A.%20Patterson%20%26%20John%20L.%20Hennessy.pdf) by David A. Patterson and John L. Hennessy  
    This is a foundational book for understanding computer architecture and how hardware and software interact. It covers instruction sets, memory, and input/output systems, all essential for designing and building a CPU.

#### 3. **Electronics Fundamentals**

- **Basic electronics knowledge**: Understanding basic principles like voltage, current, resistance (Ohm’s Law), and the role of transistors, capacitors, and resistors is necessary for working with physical components and circuits.

  - **Book**: [**The Art of Electronics**](https://eclass.uniwa.gr/modules/document/index.php?course=EEE265&download=/6218412aGLbn/5e6db8cd5KNP.pdf) by Paul Horowitz and Winfield Hill  
    This is the most comprehensive book for beginners in electronics. It covers the theory behind building and troubleshooting circuits, an essential skill for designing the hardware of your 8-bit computer.

#### 4. **Assembly Language Programming**

- **Assembly language**: You'll need to program the CPU using assembly language. This gives you direct control over the hardware, which is crucial when programming an 8-bit computer.

  - **Book**: [**Programming the Z80**](http://www.z80.info/zip/programming_the_z80_3rd_edition.pdf) by Rodnay Zaks  
    This book is a definitive guide to Z80 assembly language programming, which is representative of many 8-bit CPUs. Understanding assembly language is crucial because it’s the bridge between software and hardware.

---

## **Detailed Resources for Building an 8-bit PC**

Once you’ve covered the prerequisites, you can start the actual design and construction of an 8-bit computer. The following sections guide you through the entire process, from choosing the CPU to building memory and input/output systems.

#### 1. **Choosing the CPU or Designing Your Own**

- **CPU selection**: You can either use an existing CPU (like the Z80 or 6502) or design your own using logic gates.

  - **Resource**: [**Ben Eater’s 8-bit Breadboard Computer Series**](https://eater.net/8bit)  
    Ben Eater's series on YouTube is a hands-on, step-by-step guide to building a working 8-bit CPU on a breadboard, covering the clock, ALU, and control logic.

  - **Forum**: [**6502.org**](http://www.6502.org/)  
    This is a vibrant community focused on building computers using the 6502 processor, with tutorials, datasheets, and programming tips to help you design and build a 6502-based computer.

#### 2. **Designing and Implementing the Bus and Memory Architecture**

- **Memory mapping and bus design**: The bus allows communication between the CPU and memory. Designing the memory system includes addressing, mapping, and interfacing with the CPU.

  - **Memory Tutorial**: [**Memory Mapping in an8-bit System**](https://projects.drogon.net/6502/memory-mapping/)  
    This tutorial explains how memory addressing and mapping works in an 8-bit system.

  - **Book**: [**The Z80 Microprocessor: Architecture, Interfacing, Programming and Design**](http://www.primrosebank.net/computers/mtx/projects/mtxplus/data/books/TheZ80Microprocessor.pdf) by Ramesh Gaonkar  
    This book is essential for anyone building a Z80-based system. It covers memory and I/O interfacing techniques.

#### 3. **Building and Designing the Circuit (Breadboard or PCB)**

- **Breadboard circuits**: If you're starting with a breadboard, you'll need guidance on how to wire components together correctly.

  - **Guide**: [**Wiring and Building an 8-bit CPU on Breadboards**](https://projects.nikolasdesign.com/microcomputer-on-a-breadboard/)  
    This practical guide walks you through the wiring of an 8-bit CPU on a breadboard, including clock, memory, and buses.

  - **Tool**: [**KiCad**](https://kicad.org/)  
    An open-source PCB design tool that allows you to design printed circuit boards (PCBs) if you want to move from breadboarding to a more permanent design.

#### 4. **Programming the BIOS, Bootloader, and Writing Assembly Programs**

- **BIOS and Bootloader**: A basic BIOS (Basic Input/Output System) initializes your computer’s hardware and loads the first program into memory.

  - **Resource**: [**Writing a Simple BIOS for 8-bit Computers**](https://www.cs.bham.ac.uk/~exr/lectures/opsys/10_11/lectures/os-dev.pdf)  
    This guide explains how to write a simple BIOS that initializes the hardware and loads programs.

  - **Assembler**: [**TASM Assembler**](http://archives.oldskool.org/pub/misc/Software/Programming/tasm/Mastering%20Turbo%20Assembler/mastering_turbo_assembler_-_second_edition.pdf)  
    A simple assembler that helps you convert assembly code into machine code for your 8-bit computer.

#### 5. **Building Input/Output (I/O) Systems**

- **I/O systems**: Input/output interfaces allow communication between the computer and external devices such as keyboards or displays.

  - **Project**: [**VGA Output from an 8-bit CPU**](https://eater.net/vga)  
    This tutorial explains how to add VGA video output to your custom 8-bit computer.

  - **Book**: [**Interfacing to the Z80**](https://www.zilog.com/docs/z80/um0080.pdf)  
    This comprehensive resource covers interfacing peripherals to the Z80 processor, which applies to most 8-bit systems.

### **Additional**
#### 1. **Ben Eater’s 8-bit Computer Series**

- **YouTube Series**: [**Building an 8-bit CPU from Scratch**](https://www.youtube.com/watch?v=HyznrdDSSGM)  
  Ben Eater’s YouTube series is highly recommended for anyone interested in building an 8-bit computer from scratch. He covers every step, from the clock and ALU to memory and I/O systems.

#### 2. **Making a Retro 8-bit Computer with the Z80**

- **Blog**: [**Z80 Computer Build Blog**](https://pdfroom.com/books/build-your-own-z80-computer-steve-ciarcia/X6234b3654Z/download)  
  A step-by-step blog documenting the process of building a Z80-based 8-bit computer, including programming the ROM and interfacing with external devices.
