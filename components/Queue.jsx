import React from 'react';
import jQuery from 'jquery';

export default class Queue extends React.Component {
    getInitialState() {
        return {
            error: null,
            problemList: [],
            selectedItemIndex: -1,
            createMode: false,
            newProblem: null,
        };
    }

    constructor(props) {
        super(props);
        this.state = this.getInitialState();
    }
    
    dataToState(result) {
        if (result.error !== 0) {
            return {
                error: JSON.stringify(result.error),
            };
        }
        return {
            error: null,
            problemList: result.problem_list,
        };
    }

    loadData(cb) {
        this.setState({ error: 'Loading...' });
        if (this.serverRequest) {
            this.serverRequest.abort();
        }
        this.serverRequest = jQuery.post('/queue/list', { user: this.props.user })
                                    .done((result) => {
                                        this.setState(this.dataToState(result));
                                        console.log(result);
                                        if (cb) {
                                            cb(result);
                                        }
                                    });
                                    // .error((jqXHR, textStatus, error) => {
                                    //     this.setState({error: 'Failed to fatch data: (' + jqXHR.status + ') ' + textStatus + ' ' + error});
                                    // });
    }

    componentDidMount() {
        this.loadData();
    }

    handleClickList(index) {
        this.setState({ selectedItemIndex: index, createMode: false, newProblem: null });
    }

    handleNew() {
        const defaultNewProblem = {
            title: "",
            level: 1,
            url: "",
            status: "NOT_SOLVED",
            memo: "",
        };
        this.setState({ createMode: true, newProblem: defaultNewProblem });
    }

    renderBadgeByLevel(level) {
        if (level === 1) {
            return (<span class="badge badge-success">쉬움</span>);
        } else if (level === 2) {
            return (<span class="badge badge-warning">보통</span>)
        } else if (level === 3) {
            return (<span class="badge badge-danger">어려움</span>);
        }
    }

    renderListEntry(selectedItemIndex, elem, index) {
        const isActive = selectedItemIndex === index;
        return (
            <li
                className={isActive ? "list-group-item active" : "list-group-item"}
                onClick={this.handleClickList.bind(this, index)}
            >
                {elem.title}
                {this.renderBadgeByLevel(elem.level)}
            </li>
        )
    }

    renderList() {
        return (
            <ul className="list-group">
                {this.state.problemList.map(this.renderListEntry.bind(this, this.state.selectedItemIndex))}
            </ul>
        );
    }

    renderProblem(elem) {
        if (!elem) {
            return (<div/>);
        }

        return (
            <div className="card">
                <div class="card-body">
                    {this.renderBadgeByLevel(elem.level)}
                    <h5 class="card-title">{elem.title}</h5>
                    <a href={elem.url} class="btn btn-primary" style={{marginRight: 10}}>Problem Link</a>
                    <p class="card-text">{elem.memo}</p>
                    <a href="#" class="btn btn-info" style={{marginRight: 10}}>Edit</a>
                    <a href="#" class="btn btn-success" style={{marginRight: 10}}>Solved!</a>
                    <a href="#" class="btn btn-dark" style={{marginRight: 10}}>Push to Back</a>
                </div>
            </div>
        )
    }

    renderNewProblem() {
        return (
            <div className="card">
                <div class="card-body">
                    <span>Level</span><input type="text"></input><br/>
                    <span>Title</span><input type="text"></input><br/>
                    <span>Memo</span><input type="text"></input><br/>
                    <span>URL</span><input type="text"></input><br/>
                </div>
            </div>
        )
    }

    renderEntry() {
        if (!this.state.createMode) {
            return (
                <div>
                    {this.renderProblem(this.state.problemList[this.state.selectedItemIndex])}
                </div>
            );
        } else {
            return (
                <div>
                    {this.renderNewProblem()}
                </div>
            );
        }
    }

    render() {
        if (this.state.error != null) {
            return (
                <div> {this.state.error} </div>
            );
        }
        return (
        <div>
            <div className="row" style={{margin: 10}}>
                <button class="btn btn-primary" style={{marginRight: 10}} onClick={this.handleNew.bind(this)}>Enqueue Problem</button>
            </div>
            <div className="row" style={{margin: 20}}>
                <div className="col-md-6">
                    {this.renderList()}
                </div>
                <div className="col-md-6">
                    {this.renderEntry()}
                </div>
            </div>
        </div>
        );
    }
}

Queue.defaultProps = {
    user: null
};
