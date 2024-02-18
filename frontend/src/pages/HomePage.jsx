import { Container, Row, Col} from 'reactstrap';

/**
 * @description This component renders a HomePage to the user.
 *
 * @param None.
 * @returns {HomePage} A React element that renders a HomePage.
 */
function HomePage() {
  return (
    <>
    <Container className="mt-5"> 
      <Row className="justify-content-center"> 
      <Col xs={12} md={8}>
      <h1 className="text-center">OpenData.Swiss</h1>
      <p className="text-center">Authenticated Web Service</p>
      </Col>
      </Row>
    </Container>
    </>
  );
}

export default HomePage;