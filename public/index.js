var ReactTable = window.ReactTable.default

const columns = [{
    Header: 'User name',
    accessor: 'username'
}, {
    Header: 'Score',
    accessor: 'score'
}, {
    Header: 'Created At',
    accessor: 'created_at'
}, {
    Header: 'Updated at',
    accessor: 'updated_at'
}]

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            data: []
        };

        this.serverRequest = this.serverRequest.bind(this);
    }

    serverRequest() {
        $.get("http://localhost:3000/api/v1/scores", res => {
            this.setState({
                data: res
            });
        });
    }

    componentDidMount() {
        this.serverRequest();
    }

    render() {
        return <ReactTable defaultPageSize={10} showPagination={false} data={this.state.data} columns={columns} />
    }
}

ReactDOM.render(<App />, document.getElementById("root"));
