import swal from 'sweetalert'
import {loadingWrapper, setData} from './common'

const unknownUser = {unknown: true}
const setUser = setData(['loggedUser'])

export const getMyInfo = () => (
  async (dispatch, getState, {api}) => {
    dispatch(setUser({loading: true}))
    if (localStorage.getItem('token')) {
      try {
        const userData = await api.getUserInfo()
        dispatch(setUser(userData))
      } catch (error) {
        localStorage.removeItem('token')
        dispatch(setUser(unknownUser))
      }
    } else {
      dispatch(setUser(unknownUser))
    }
  }
)

export const updateUser = (data) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const userData = await api.updateUser(data)
      dispatch(setUser(userData))
    } catch (error) {
      await swal({
        title: 'User data could not be updated',
        text: error.message,
        icon: 'error',
      })
    }
  }
)

export const loginWithSlovenskoSkToken = (token) => (
  async (dispatch, getState, {api}) => {
    try {
      const userData = await api.loginWithSlovenskoSkToken(token)
      dispatch(setUser(userData))
      localStorage.setItem('token', userData.token)
      return true
    } catch (error) {
      dispatch(setUser(unknownUser))
      localStorage.removeItem('token')
      await swal({
        title: 'Login failed',
        text: error.message,
        icon: 'error',
      })
      return false
    }
  }
)

export const logout = () => (
  async (dispatch, getState, {api}) => {
    await api.logout()
    dispatch(setUser(unknownUser))
    localStorage.removeItem('token')
  }
)
