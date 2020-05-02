
const limit = 1000;

/**
 * @interface
 */
const Request = {
  ID: Number,
  RawHeaders: String,
  RawBody: String,
  Method: String,
  URI: String,
  StatusCode: Number
}

/**
 * @interface
 */
const defaultState = {
  $tabPage: 'HTTP' | 'TCP' | 'DRIVER',
  /**
   * @type {Request[]}
   */
  $logs: [],
  /**
   * @type {Request[]}
   */
  $filtered_logs: [],
  $search_text: String,
  $request: Request || null,
  $page: Number,
  $lastPage: 1
}

const store = {
  /**
   * @type {defaultState}
   */
  state: {
    $tabPage: 'HTTP',
    $logs: [],
    $filtered_logs: [],
    $search_text: '',
    $request: null,
    $page: 1,
    $lastPage: 1
  },

  /**
   * 
   * @param {defaultState} state 
   */
  setState(state) {
    Object.assign(this.state, state)
  },

  /**
   * 
   * @param {'HTTP'|'TCP'|'DRIVER'} tabPage 
   */
  setTabPage(tabPage) {
    this.setState({ $tabPage: tabPage, $page: 1 })
  },

  /**
   * 
   * @param {Number} total 
   */
  lastPage(total) {
    let lastPage = parseInt(total / limit);
    if (lastPage * limit !== total) lastPage += 1
    if (lastPage === 0) lastPage = 1;
    this.setState({ $lastPage: lastPage })
  },

  /**
   * 
   * @param {String|Request} request
   */
  async setRequest(request) {
    if (typeof request === 'object') {
      this.setState({
        $request: request
      })
      return
    }
    if (!request) this.setState({ $request: null })
    await get_request(request)
  },

  /**
   * 
   * @param {Number} page 
   */
  setPage(page) {
    this.setState({
      $page: parseInt(page),
      $filtered_logs: this._getLogsPage(page, this.state.$logs)
    })
  },

  _getLogsPage(page, logs) {
    return logs.slice((page - 1) * limit, page * limit)
  },

  setSearchText(search_text) {
    this.setState({
      $search_text: search_text
    })
  },

  filter() {
    if (this.state.$search_text === '') return;
    const logs = this.state.$logs.filter(() => {
      let preg = new RegExp(this.state.$search_text)
      if (preg.test(this.state.$request.RawRequest) ||
        preg.test(this.state.$request.RawResponse) ||
        preg.test(this.state.$request.URI) ||
        preg.test(this.state.$request.StatusCode.toString()) ||
        preg.test(this.state.$request.Method) ||
        preg.test(this.state.$request.Time))
        return true
      return false
    })
    this.setState({
      $filtered_logs: this._getLogsPage(this.state.$page, logs)
    })
  },

  /**
   * 
   * @param {Request[]} logs 
   */
  setLogs(logs = []) {
    this.setState({
      $logs: logs,
      $filtered_logs: this._getLogsPage(this.state.$page, logs)
    })
    this.filter()
  },

  mapColor() {
    const logs = this.state.$logs.map((el) => {
      if (el.StatusCode >= 200 && el.StatusCode < 300)
        return { ...el, color: 'green' }
      if (el.StatusCode >= 300 && el.StatusCode < 400)
        return { ...el, color: 'orange' }
      if (el.StatusCode >= 400 && el.StatusCode < 500)
        return { ...el, color: 'purple' }
      if (el.StatusCode >= 500 && el.StatusCode < 600)
        return { ...el, color: 'red' }
      if (el.StatusCode >= 100 && el.StatusCode < 200)
        return { ...el, color: 'blue' }
      return { ...el, color: 'black' }
    })
    this.setState({
      $filtered_logs: this._getLogsPage(this.state.$page, logs)
    })
  }
}