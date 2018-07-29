import React from 'react';
import Queue from './Queue.jsx';

class App extends React.Component {
    render(){

        return (
            <div>
                <h1>Hello PSQueue</h1>
                <br/>
                <Queue
                    user={"ethanhur"}
                />
            </div>
        );
    }
}

export default App;