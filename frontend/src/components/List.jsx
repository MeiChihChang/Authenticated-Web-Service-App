import { Table, Row, Col } from "reactstrap"

/**
 * @description This component renders a datsset table to the user.
 *
 * @param {[data]} datalist dataset list.
 * @returns {DataList} A React element that renders a datsset list table by selected organization name.
 */
function DataList({datalist}) {
  //console.log("datalist", datalist.length, datalist)
  return (

      <Row className="justify-content-center"> 
      <Col>
        <h1>All Data</h1>
        
        <Table bordered={true}>
            <thead> 
                <tr> 
                    <th className={"w-10"}>Owner_org</th> 
                    <th className={"w-10"}>Maintainer</th> 
                    <th className={"w-10"}>Issued</th> 
                    <th className={"w-10"}>Maintainer_email</th> 
                    <th className={"w-60"}>download_url</th>
                </tr> 
            </thead> 
            <tbody>
                {datalist.map((data) => (
                <tr key={data.id} >
                    <td>{data.owner_org}</td>
                    <td>{data.maintainer}</td>
                    <td>{data.issued}</td>
                    <td>{data.maintainer_email}</td>
                    <td><a href={data.download_url}>{data.download_url}</a></td>
                </tr>
                ))}
            </tbody> 
        </Table>
        
      </Col>
      </Row>
  );
}

export default DataList;