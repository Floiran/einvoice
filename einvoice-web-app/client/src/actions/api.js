const setApiInitialized = () => ({
  type: 'SET API INITIALIZED',
  path: ['apiInitialized'],
  payload: null,
  reducer: (state) => true,
})

// This should not be necessary in future as API url can be send in initial HTML
export const initializeApi = () => (
  async (dispatch, getState, {api}) => {
    await api.getApiUrl()
    dispatch(setApiInitialized())
  }
)
