import Header from './components/Header'
import Events from './components/Events'
import AddEvent from './components/AddEvent'
import { useState } from 'react'

function App() {
  const [showAddEvent, setShowAddEvent] = useState(false)
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

  // Add Event
  const addEvent = (singleEvent) => {
    // TODO: have this call the backend POST   /api/v1/secured/event
    const id = Math.floor(Math.random() * 10000) + 1
    const newEvent = { id, ...singleEvent }
    console.log(newEvent)
    setEvents([...events, newEvent])
  }

  // Delete Event
  const deleteEvent = (id) => {
    // TODO: have this call the backend DELETE /api/v1/secured/event/:id
    setEvents(events.filter((event) => event.id !== id))
  }

  return (
    <div className="container">
      <Header title='My Events'
        canAddEvent={true}
        showAdd={showAddEvent}
        onAdd={() => setShowAddEvent(!showAddEvent)} />
      {showAddEvent && <AddEvent onAdd={addEvent} />}
      {events.length > 0 ? (<Events events={events} includeDeleteButton={true} onDelete={deleteEvent} />)
        : ('Create an event!')}
      <br />
      <br />
      <Header title='Joined Events' />
      {events.length > 0 ? (<Events events={events} />)
        : ('Join an event!')}
    </div>
  );
}

export default App;
