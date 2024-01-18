import { BigNumberish } from 'ethers';

export function splitHash(h: string): BigNumberish[] {
  if (h.startsWith('0x')) {
    h = h.slice(2);
  }
  const a = '0x' + h.substring(0, h.length / 2);
  const b = '0x' + h.substring(h.length / 2, h.length);
  return [a, b];
}

export function hexToBytes(hex: string) {
  if (hex.startsWith('0x')) {
    hex = hex.slice(2);
  }
  for (var bytes = [], c = 0; c < hex.length; c += 2) bytes.push(parseInt(hex.slice(c, c + 2), 16));
  return bytes;
}
