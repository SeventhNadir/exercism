//
// This is only a SKELETON file for the 'Pangram' exercise. It's been provided as a
// convenience to get you started writing code faster.
//
const lowercaseLetters = "abcdefghijklmnopqrstuvwxyz";

export const isPangram = (candidatePangram) => {
  candidatePangram = candidatePangram.toLowerCase();
  
  return [...lowercaseLetters].every(char => candidatePangram.includes(char))

};
