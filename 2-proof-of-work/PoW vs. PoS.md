## Proof of Work vs. Proof of Stake

At the heart of every cryptocurrency lies a network of computers that helps secure the software from attackers and regulates the issuance of new units of its supply. 

This system is called a *consensus mechanism*.

The two most widely used consensus mechanisms are Proof of Work (PoW) and Proof of Stake (PoS), and they both regulate the process in which transactions between users are verified and added to a blockchain’s public ledger, all without a central party’s help.



#### ***What is Proof of Work (PoW)?***

Proof of Work (PoW) was introduced in the early 1990s as a means to mitigate email spam. 

The idea was computers might be required to perform a small amount of work before sending an email. This work would be trivial for someone sending a legitimate email, but it would require a lot of computing power and resources for users to send mass emails.

However, it was Satoshi Nakamoto, Bitcoin’s creator, who first applied the technology for use in a digital money system in Bitcoin’s white paper.

##### Blockchain Ordering

A blockchain is a mechanism that consists of a string of blocks (groups of transactions) aligned in chronological order based on the order of transactions.

The first block in a PoW blockchain is hardcoded into its software and is named the genesis block, also known as block 0. By definition, this block does not reference a previous block. 

The subsequent blocks added to the blockchain always refer to the previous blocks and each contain a copy of the full, updated ledger.

##### Energy Use

PoW algorithms determine who can adjust the ledger through a competitive race in which certain participants (miners) are encouraged to expend computational energy in order to propose valid blocks that meet the rules of the network.

Nodes (any computer running the Bitcoin software) then validate transactions, prevent double spending (the act of spending the same funds to two separate recipients) and determine whether proposed blocks should be added to the chain.

In order to create a new block, miners on a PoW network compete against each other to solve complex mathematical problems in a process called hashing. These puzzles are very hard to solve, but should be easy for the network to verify the correct solution. 

##### Participation

In a PoW protocol, computational power is combined with cryptography to create consensus and ensure the validity of transactions recorded in the blockchain. 

During the hashing process, and in order to produce new blocks, miners race to generate a correct result to the mathematical problems.

In order to do so, miners guess a string of pseudorandom numbers, called a hash. This, when combined with the data provided in the block, and when passed through a hash function computer, must produce a result that matches the given conditions set out by the protocol. 

The winning hash is subsequently broadcasted to the network for the other miners to verify whether the solution is true or not. If it is correct, the block gets added to the blockchain and the miner gets compensated with the block reward.

##### Reward Distribution 

The block reward refers to new cryptocurrency awarded by the blockchain to the miner for each block that is deemed valid and accepted by the network.

In the case of certain cryptocurrencies, like Bitcoin, the block reward is reduced after a certain number of blocks have been found. 

This is done to keep the total money supply finite and deflationary.



#### ***What is Proof of Stake (PoS)***

Proof of Stake (PoS) is a modification of PoW introduced in 2012 as a means to solve its perceived dependency on energy consumption as a means to determine blockchain ordering.

Rather than rely on computers racing to generate the appropriate hash, the idea behind a PoS protocol is that participation is determined by ownership of the coin supply.

Using a set of factors determined by the protocol, the PoS algorithm pseudo-randomly elects a node (anyone who owns the coin) to propose the next block to the blockchain. 

When a node gets elected, its role is to verify the validity of the transactions within the block, sign it and propose the block to the network for validation.

##### Blockchain Ordering

Similar to PoW, a PoS blockchain is a mechanism that consists of a string of blocks aligned in chronological order based on its transactions.

The first block in a PoS blockchain is also hardcoded into its software and is commonly named the genesis block. The subsequent blocks added to the blockchain always refer to the previous blocks and each contain a copy of the full, updated ledger.

Of note, in PoS cryptocurrencies, there is no competition for who is selected to add blocks. As such, the blocks are often said to be ‘forged’, or ‘minted’ rather than mined.

##### Energy Use

Unlike PoW blockchains, PoS blockchains do not determine who can propose blocks based on energy consumption. Proponents often bill PoS as a “more efficient” system, as PoW miners have been known to use coal and other forms of fuel harmful to the environment.

That said, PoW miners can choose to use any energy source, including hydroelectric power, wind and other less wasteful forms of energy. 

Further, PoS blockchains also require the use of specialized hardware (GPUs) that, like PoW mining equipment (ASICs) and other computing devices, require resources to produce. 

PoS miners must also maintain active internet connections, which requires energy expense.

##### Participation

Users who want the opportunity to be selected to add blocks to a PoS blockchain are required to stake, or lock, a certain amount of the blockchain’s cryptocurrency in a special contract. 

The amount of coins staked determines their chances to be selected as the next block producer. In some cases, if users behave maliciously, they may lose their stake as punishment. 

In order to not always favor the wealthiest nodes, PoS may involve other determining factors. These may include the amount of time a node has staked their coins, as well as pure randomization. 

##### Reward Distribution

Similar to the PoW algorithm, the block reward in PoS refers to cryptocurrency awarded by the blockchain to the user who proposes a valid block. 

However, as block selection is done by coin ownership, exchanges may offer staking services that offer users the ability to stake funds on their behalf in exchange for more regular payouts.