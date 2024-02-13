import { Form, useNavigation } from 'react-router-dom';

import { Container, Row, Col, Button, Input, Label} from 'reactstrap'


function AuthForm() {
  const navigation = useNavigation();
  const isSubmitting = navigation.state === 'submitting';

  return (
    <>
    <Container className="mt-5"> 
      <Row className="justify-content-center"> 
      <Col xs={12} md={8}>
      <Form method="post">
        <h1 className="text-center">Log in</h1>
        <p>
          <Label htmlFor="username">Username</Label>
          <Input id="username" type="username" name="username" required />
        </p>
        <p>
          <Label  htmlFor="password">Password</Label>
          <Input id="password" type="password" name="password" required />
        </p>
        <Button color="primary" disabled={isSubmitting} type="submit">
          {isSubmitting ? 'Submitting...' : 'Login'}
        </Button>
      </Form>
      </Col>
      </Row>
      </Container>
    </>
  );
}

export default AuthForm;