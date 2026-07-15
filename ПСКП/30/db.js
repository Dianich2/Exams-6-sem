const { Sequelize, DataTypes } = require('sequelize');

const sequelize = new Sequelize('UNIVER', 'sa', '1234', {
 host: 'localhost',
  dialect: 'mssql',
  dialectOptions: {
    options: {
      encrypt: false,
      trustServerCertificate: true
    }
  },
  pool: { max: 5, min: 0, idle: 10000 },
  define: { timestamps: false }
});

const Faculty = sequelize.define('FACULTY', {
  faculty: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  faculty_name: { type: DataTypes.STRING }
}, { tableName: 'FACULTY' });

const Pulpit = sequelize.define('PULPIT', {
  pulpit: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  pulpit_name: { type: DataTypes.STRING },
  faculty: { type: DataTypes.STRING }
}, { tableName: 'PULPIT' });

const Subject = sequelize.define('SUBJECT', {
  subject: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  subject_name: { type: DataTypes.STRING },
  pulpit: { type: DataTypes.STRING }
}, { tableName: 'SUBJECT' });

const Teacher = sequelize.define('TEACHER', {
  teacher: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  teacher_name: { type: DataTypes.STRING },
  gender: { type: DataTypes.STRING },
  pulpit: { type: DataTypes.STRING }
}, { tableName: 'TEACHER' });

const Auditorium = sequelize.define('AUDITORIUM', {
  auditorium: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  auditorium_type: { type: DataTypes.STRING },
  auditorium_capacity: { type: DataTypes.INTEGER },
  auditorium_name: { type: DataTypes.STRING }
}, { tableName: 'AUDITORIUM' });

const AuditoriumType = sequelize.define('AUDITORIUM_TYPE', {
  auditorium_type: { type: DataTypes.STRING, primaryKey: true, allowNull: false },
  auditorium_typename: { type: DataTypes.STRING }
}, { tableName: 'AUDITORIUM_TYPE' });

Faculty.hasMany(Pulpit, {
  foreignKey: 'faculty',
  sourceKey: 'faculty'
});

Pulpit.belongsTo(Faculty, {
  foreignKey: 'faculty',
  targetKey: 'faculty'
});

Pulpit.hasMany(Subject, {
  foreignKey: 'pulpit',
  sourceKey: 'pulpit'
});

Subject.belongsTo(Pulpit, {
  foreignKey: 'pulpit',
  targetKey: 'pulpit'
});

Pulpit.hasMany(Teacher, {
  foreignKey: 'pulpit',
  sourceKey: 'pulpit'
});

Teacher.belongsTo(Pulpit, {
  foreignKey: 'pulpit',
  targetKey: 'pulpit'
});


module.exports = {
    sequelize,
    Faculty,
    Pulpit,
    Subject,
    Teacher,
    Auditorium,
    AuditoriumType
};