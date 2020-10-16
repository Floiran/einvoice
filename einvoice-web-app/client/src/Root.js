import React from 'react'
import {Provider} from 'react-redux'
import {BrowserRouter} from 'react-router-dom'
import App from './components/App'
import ErrorBoundary from './components/helpers/ErrorBoundary'

const Spinner = () => (
  <div className="Modal">
    <div className="loader" />
  </div>
)

export default ({store}) => (
  <ErrorBoundary>
    <Provider store={store}>
      <BrowserRouter>
        <React.Suspense fallback={<Spinner />}>
          <App />
        </React.Suspense>
      </BrowserRouter>
    </Provider>
  </ErrorBoundary>
)
