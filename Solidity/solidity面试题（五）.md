## **gas优化**

gas优化有哪些手段，也是面试中常被问到的。这里我不详细给出123这样的列表。大家可以参考这个[gas优化手册](https://decert.me/tutorial/rareskills-gas-optimization/chapter_1)。

在evm中，对存储的操作都是要消耗gas的，例如读取，写入，计算等。不同的操作对应有不同的操作码，每个操作码对应的gas也是不一样的。

为了节省gas，我们可以减少对状态变量的操作，例如利用常量constant和immutable修饰不修改的变量，原因大家在[gas优化手册](https://decert.me/tutorial/rareskills-gas-optimization/chapter_1)这里找。

尽量使用少的slot存储槽。例如，uint80 ，uint80，uint256，顺序只用了两个存储槽，而uint80 , uint256 , uint80 使用了3个存储槽。

在前面的面试题有些问题也直接或间接的涉及到gas的，也可以当做这类问题来回答。

## 部分Solidity操作码的gas列表：

| **操作码** | **操作**       | **Gas消耗** |
| :--------- | :------------- | :---------- |
| 0x00       | STOP           | 0           |
| 0x01       | ADD            | 3           |
| 0x02       | MUL            | 5           |
| 0x03       | SUB            | 3           |
| 0x04       | DIV            | 5           |
| 0x05       | SDIV           | 5           |
| 0x06       | MOD            | 5           |
| 0x07       | SMOD           | 5           |
| 0x08       | ADDMOD         | 8           |
| 0x09       | MULMOD         | 8           |
| 0x0a       | EXP            | 20          |
| 0x0b       | SIGNEXTEND     | 5           |
| 0x10       | LT             | 3           |
| 0x11       | GT             | 3           |
| 0x12       | SLT            | 3           |
| 0x13       | SGT            | 3           |
| 0x14       | EQ             | 3           |
| 0x15       | ISZERO         | 3           |
| 0x16       | AND            | 3           |
| 0x17       | OR             | 3           |
| 0x18       | XOR            | 3           |
| 0x19       | NOT            | 3           |
| 0x1a       | BYTE           | 3           |
| 0x20       | SHA3           | 30          |
| 0x30       | ADDRESS        | 2           |
| 0x31       | BALANCE        | 400         |
| 0x32       | ORIGIN         | 2           |
| 0x33       | CALLER         | 2           |
| 0x34       | CALLVALUE      | 2           |
| 0x35       | CALLDATALOAD   | 3           |
| 0x36       | CALLDATASIZE   | 2           |
| 0x37       | CALLDATACOPY   | 3           |
| 0x38       | CODESIZE       | 2           |
| 0x39       | CODECOPY       | 3           |
| 0x3a       | GASPRICE       | 2           |
| 0x3b       | EXTCODESIZE    | 700         |
| 0x3c       | EXTCODECOPY    | 700         |
| 0x3d       | RETURNDATASIZE | 2           |
| 0x3e       | RETURNDATACOPY | 3           |
| 0x40       | BLOCKHASH      | 20          |
| 0x41       | COINBASE       | 2           |
| 0x42       | TIMESTAMP      | 2           |
| 0x43       | NUMBER         | 2           |
| 0x44       | DIFFICULTY     | 2           |
| 0x45       | GASLIMIT       | 2           |
| 0x50       | POP            | 2           |
| 0x51       | MLOAD          | 3           |
| 0x52       | MSTORE         | 3           |
| 0x53       | MSTORE8        | 3           |
| 0x54       | SLOAD          | 200         |
| 0x55       | SSTORE         | 20000       |
| 0x56       | JUMP           | 8           |
| 0x57       | JUMPI          | 10          |
| 0x58       | PC             | 2           |
| 0x59       | MSIZE          | 2           |
| 0x5a       | GAS            | 2           |
| 0x5b       | JUMPDEST       | 1           |
| 0xa0       | LOG0           | 375         |
| 0xa1       | LOG1           | 750         |
| 0xa2       | LOG2           | 1125        |
| 0xa3       | LOG3           | 1500        |
| 0xa4       | LOG4           | 1875        |
| 0xf0       | CREATE         | 32000       |
| 0xf1       | CALL           | 700         |
| 0xf2       | CALLCODE       | 700         |
| 0xf3       | RETURN         | 0           |
| 0xf4       | DELEGATECALL   | 700         |
| 0xfa       | STATICCALL     | 700         |
| 0xfd       | REVERT         | 0           |
| 0xff       | SELFDESTRUCT   | 5000        |
