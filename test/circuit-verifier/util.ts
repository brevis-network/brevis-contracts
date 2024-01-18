export const convertByteArrayToHexString = (input: number[]) => {
  var result = '';
  input.forEach((value) => {
    var hexString = value.toString(16);
    if (hexString.length % 2 == 1) {
      hexString = '0' + hexString;
    }
    result += hexString;
  });

  return '0x' + result;
};
