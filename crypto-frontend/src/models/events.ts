export interface PayPerViewEventProspect {
  id?: number;
  name: string;
  description: string;
  eventType: string;
  start: Date;
  end: Date;
  priceBTC: number;
  priceXMR: number;
  priceETH: number;
}
