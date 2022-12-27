export interface ICreatureTypeOptions {
  name?: string;
  str?: number;
  dex?: number;
  wis?: number;
  con?: number;
  int?: number;
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
  senses?: {
    darkvision?: number;
    passive?: number;
  };
  type?: string;
  align?: string;
  features?: string[];
  actions?: {
    name: string;
    desc: string;
  }[];
  reactions?: {
    name: string;
    desc: string;
  }[];
  other?: object;
}

export class CreatureType {
  constructor() {}
}
