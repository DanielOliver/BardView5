import { z } from 'zod'
import { CommonAccessValues } from '../common'

const Dnd5eSettingCreateSchema = z.object({
  name: z.string().min(1).max(512),
  description: z.string().min(1).max(1024),
  module: z.string().optional(),
  commonAccess: z.enum(CommonAccessValues).default('private'),
  userTags: z.string().array().default([]),
  systemTags: z.string().array().default([]),
  active: z.boolean().default(true)
})

type Dnd5eSettingCreateType = z.infer<typeof Dnd5eSettingCreateSchema>

export type {
  Dnd5eSettingCreateType
}

export {
  Dnd5eSettingCreateSchema
}
