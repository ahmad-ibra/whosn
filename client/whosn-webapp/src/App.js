import Header from './components/Header'
import Events from './components/Events'
import { useState } from 'react'

function App() {
  const [events, setEvents] = useState([
    {
      id: 1,
      text: 'Meeting at school',
      day: 'Feb 6th at 1:30pm',
      reminder: true
    },
    {
      id: 2,
      text: 'Finish whosn',
      day: 'Feb 23rd at 8:00pm',
      reminder: true
    },
    {
      id: 3,
      text: 'bla bla bla',
      day: 'Mar 13th at 9:30am',
      reminder: false
    }
  ])

  return (
    <div className="container">
      <Header title='My Events' canAddEvent={true} />
      <Events includeDeleteButton={true} events={events} />
      <br />
      <br />
      <Header title='Joined Events' />
      <Events events={events} />
    </div>
  );
}

export default App;
