import { ComponentMeta, ComponentStory, Story } from '@storybook/react'
import { Layout } from '../components/Layout'
import React from 'react'
import { AuthContext } from '../context/Auth.context'

export default {
  title: 'Global/Layout',
  component: Layout
} as ComponentMeta<typeof Layout>

const Template: ComponentStory<typeof Layout> = () => <Layout logout={() => {}}/>

export const LoggedIn = Template.bind({ isAuthenticated: true })
LoggedIn.decorators = [
  (Story) => <AuthContext.Provider value={{
    state: {
      isAuthenticated: true,
      isRegistrationComplete: true,
      checked: 'CHECKED'
    },
    dispatch: reduce => {
    }
  }}>
    <Story/>
  </AuthContext.Provider>
]
export const LoggedOut = Template.bind({ isAuthenticated: false })
LoggedOut.decorators = [
  (Story) => <AuthContext.Provider value={{
    state: {
      isAuthenticated: false,
      isRegistrationComplete: true,
      checked: 'CHECKED'
    },
    dispatch: reduce => {
    }
  }}>
    <Story/>
  </AuthContext.Provider>
]
