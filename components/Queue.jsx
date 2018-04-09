import React from 'react';

export default class Queue extends React.Component {
    render() {
        return (
        <div>
            <div className="row" style={{margin: 10}}>
                <a href="#" class="btn btn-primary" style={{marginRight: 10}}>Add New</a>
            </div>
            <div className="row" style={{margin: 20}}>
                <div className="col-md-6">
                    <ul className="list-group">
                        <li className="list-group-item active">별찍기1<span class="badge badge-danger">어려움</span></li>
                        <li className="list-group-item">별찍기2<span class="badge badge-success">쉬움</span></li>
                        <li className="list-group-item">별찍기3</li>
                    </ul>
                </div>
                <div className="col-md-6">
                    <div className="card">
                        <div class="card-body">
                            <span class="badge badge-danger">어려움</span>
                            <h5 class="card-title">별찍기1</h5>
                            <p class="card-text">류가 내준 문제 쏼라쏼라 궁시렁궁시렁</p>
                            <a href="#" class="btn btn-primary" style={{marginRight: 10}}>Problem Link</a>
                            <a href="#" class="btn btn-info" style={{marginRight: 10}}>Edit</a>
                            <a href="#" class="btn btn-success" style={{marginRight: 10}}>Solved!</a>
                            <a href="#" class="btn btn-dark" style={{marginRight: 10}}>Push to Back</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        );
    }
}