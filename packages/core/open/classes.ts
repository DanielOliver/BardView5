export interface Bv5 {
  id: string;
  t: string;
  name: string;
}

export interface Bv5Session {
  id: string;
  name: string;
  date?: string;
}

export interface ICreatureTypeOptions {
  name: string;
  str?: number;
  dex?: number;
  con?: number;
  int?: number;
  wis?: number;
  cha?: number;
  walking?: number;
  flying?: number;
  climbing?: number;
  prof?: string[];
  bonus?: number;
  equip?: string[];
  lang?: string[] | string;
  cr?: number;
  xp?: number;
  hpFlat?: number;
  hpRoll?: string;
  ac?: number;
  throws?: string[];
  skills?: {
    name: string;
    bonus: number;
  }[];
  darkvision?: number;
  passive?: number;
  type?: string;
  align?: string;
  features?: string[];
  actions?: string[];
  reactions?: {
    name: string;
    desc: string;
  }[];
  other?: object;
}

export interface Bv5Session {
  id: string;
  name: string;
  date?: string;
}
