export const SATOSHI = 100000000;
export const PICONERO = 1000000000000;
export const WEI = 1000000000000000000;

export const BTC = "btc";
export const XMR = "xmr";
export const ETH = "eth";

interface CryptonatorTicker {
  base: string;
  target: string;
  price: string;
  volume: string;
  change: string;
}

interface CryptonatorResponse {
  ticker: CryptonatorTicker;
  timestamp: number;
  success: boolean;
  error: string;
}

export interface CryptoCurrencyValues {
  satoshi: number;
  piconero: number;
  wei: number;
}

export const btcToSatoshi = function(btc: number): number {
  return Math.floor(btc * SATOSHI);
};

export const xmrToPiconero = function(xmr: number): number {
  return Math.floor(xmr * PICONERO);
};

export const ethToWei = function(eth: number): number {
  return Math.floor(eth * WEI);
};

export const usdToBTC = function(usd: number) {
  return usdToCryptoCurrency(usd, BTC, btcToSatoshi);
};

export const usdToXMR = function(usd: number) {
  return usdToCryptoCurrency(usd, XMR, xmrToPiconero);
};

export const usdToETH = function(usd: number) {
  return usdToCryptoCurrency(usd, ETH, ethToWei);
};

export const usdToCrypto = async function(
  usd: number
): Promise<CryptoCurrencyValues> {
  const currencyPrice = await Promise.all([
    usdToBTC(usd),
    usdToXMR(usd),
    usdToETH(usd)
  ]);
  const [satoshi, piconero, wei] = currencyPrice;
  return {
    satoshi,
    piconero,
    wei
  };
};

const usdToCryptoCurrency = async function(
  usd: number,
  cryptoCurrency: string,
  valueConverter: (val: number) => number
) {
  const response = await fetch(
    `https://api.cryptonator.com/api/ticker/${encodeURIComponent(
      cryptoCurrency
    )}-usd`
  );
  const cryptonatorResponse = (await response.json()) as CryptonatorResponse;
  if (!cryptonatorResponse.success) {
    throw new Error(cryptonatorResponse.error);
  }
  const usdToCrypto = parseFloat(cryptonatorResponse.ticker.price);
  return valueConverter(usd / usdToCrypto);
};
