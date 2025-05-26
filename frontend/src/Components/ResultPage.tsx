import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Typography,
  Box,
  Button,
} from "@mui/material";
import { useLocation, useNavigate } from "react-router";

const ResultPage = () => {
  const navigate = useNavigate();

  const location = useLocation();
  const data = location.state;
  const {
    koszt,
    przychod,
    tabelaTransportu,
    zyskCalkowity,
    tabelaZyskowJednostkowych,
  } = data;

  return (
    <Box sx={{ padding: 3 }}>
      {/* Wyświetlenie kosztów, przychodu i zysku */}
      <Box sx={{ marginBottom: 2 }}>
        <Typography variant="body1">
          <strong>Koszt:</strong> {koszt} zł
        </Typography>
        <Typography variant="body1">
          <strong>Przychód:</strong> {przychod} zł
        </Typography>
        <Typography variant="body1">
          <strong>Zysk całkowity:</strong> {zyskCalkowity} zł
        </Typography>
      </Box>

      {/* Tabela transportu */}
      <Typography variant="h5" gutterBottom>
        Tabela Transportu
      </Typography>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="tabela transportu">
          <TableHead>
            <TableRow>
              {tabelaTransportu[0]?.map((_, index) => (
                <TableCell key={index} align="center">
                  Odbiorca {index + 1}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {tabelaTransportu.map((row, rowIndex) => (
              <TableRow key={rowIndex}>
                {row.map((cell, cellIndex) => (
                  <TableCell key={cellIndex} align="center">
                    {cell}
                  </TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Typography variant="h5" gutterBottom>
        Tabela Zyskow Jednostkowych
      </Typography>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="tabela transportu">
          <TableHead>
            <TableRow>
              {tabelaZyskowJednostkowych[0]?.map((_, index) => (
                <TableCell key={index} align="center">
                  Odbiorca {index + 1}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {tabelaZyskowJednostkowych.map((row, rowIndex) => (
              <TableRow key={rowIndex}>
                {row.map((cell, cellIndex) => (
                  <TableCell key={cellIndex} align="center">
                    {cell}
                  </TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Button variant="contained" onClick={() => navigate("/")}>
        Home
      </Button>
    </Box>
  );
};

export default ResultPage;
