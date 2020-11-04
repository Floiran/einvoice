export const setData = (path) => (data) => ({
  type: `SET DATA ON ${path}`,
  path,
  payload: data,
  reducer: (_, data) => data,
})

const setLoadingState = setData(['isLoading'])

export const loadingWrapper = (action) => (
  async (dispatch) => {
    dispatch(setLoadingState(true))
    const returnValue = await dispatch(action)
    dispatch(setLoadingState(false))
    return returnValue
  }
)
