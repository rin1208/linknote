import React from "react";

import Grid from "@material-ui/core/Grid";

import styled from "@emotion/styled";
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';
import { ReactTinyLink } from 'react-tiny-link'

import Typography from '@material-ui/core/Typography';


const Main = styled.div`
display: "flex";
margin: 20;
text-align: center;
`;

const Sub = styled.div`
  text-align: center;
`;

const Footer = styled.footer`
margin-top: auto;
`;



function Content() {
    var front = [
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin_kannzaki",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin1208_sugi",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin_kannzaki",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin1208_sugi",
        },

        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin_kannzaki",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin1208_sugi",
        },

        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin_kannzaki",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin1208_sugi",
        },

        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin_kannzaki",
        },
        {
            title: "Twitter",
            coment: "ほげ",
            url: "https://twitter.com/rin1208_sugi",
        },

    ];
    var list = front.map(function (item) {
        return (
            <Grid item xs={6} md={3}>

                <Card >
                    <CardContent>
                        <Typography  >
                            <ReactTinyLink
                                cardSize="small"
                                showGraphic={true}
                                maxLine={2}
                                minLine={1}
                                url={item.url}
                            />
                        </Typography>
                    </CardContent>

                    <Sub>
                        <CardActions>
                            <Button size="big" color="primary">
                                Delete
                        </Button>
                        </CardActions>
                    </Sub>

                </Card>

            </Grid>
        );
    });
    return (
        <div>


            <Main>
                <Grid container spacing={1}>
                    <Grid container item xs={12} spacing={2}>
                        {list}

                    </Grid>
                </Grid>
            </Main>

            <Main>
                <Footer>
                    <p>(c)copy right</p>
                </Footer>
            </Main>

        </div>
    );
}

export default Content;