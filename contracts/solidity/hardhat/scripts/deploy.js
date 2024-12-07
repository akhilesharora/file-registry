const hre = require("hardhat");

async function main() {
  const FileRegistry = await hre.ethers.getContractFactory("FileRegistry");
  const network = await hre.ethers.provider.getNetwork();
  console.log("Deploying to network:", network.name);
  const fileRegistry = await FileRegistry.deploy();
  // Wait for the contract to be mined
  await fileRegistry.waitForDeployment();
  console.log("FileRegistry deployed to:", await fileRegistry.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
