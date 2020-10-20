import swal from 'sweetalert'
import {setData} from './common'

// This should not be necessary in future as API urls can be send in initial HTML
export const initializeApi = () => (
  async (dispatch, getState, {api}) => {
    try {
      const urls = await api.getApiUrl()
      dispatch(setData(['urls'])(urls))
      dispatch(setData(['apiInitialized'])(true))
    } catch(error) {
      await swal({
        title: 'API urls could not be loaded',
        text: error.message,
        icon: 'error',
      })
    }
  }
)
