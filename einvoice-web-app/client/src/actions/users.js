export const setUser = (user) => ({
  type: 'SET USER',
  path: ['user'],
  payload: user,
  reducer: (state, user) => user,
})

export const getMyInfo = () => (
  async (dispatch, getState, {api}) => {
    let userString = localStorage.getItem('user')
    let user = userString && JSON.parse(userString)

    if (user) {
      try {
        let userData = await api.getUserInfo(user)
        dispatch(setUser(userData))
      } catch (error) {
        if (error.statusCode === 401) {
          dispatch(setUser(null))
        }
      }
    }
  }
)

export const login = () => (
  async (dispatch, getState, {api}) => {
    const userData = await api.login()
    dispatch(setUser(userData))
    localStorage.setItem('user', JSON.stringify(userData))
  }
)

export const logout = () => (
  async (dispatch, getState, {api}) => {
    await api.logout(getState().user)
    dispatch(setUser(null))
    localStorage.removeItem('user')
  }
)
